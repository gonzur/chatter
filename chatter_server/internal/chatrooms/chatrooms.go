package chatrooms

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var socketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return socketUpgrade.Upgrade(w, r, nil)
}

type floorMap map[string]*Room

/*
Multithreading safe access to current rooms.
*/
type safeLobby struct {
	registeredRooms floorMap
	process         chan func(floorMap)
}

func (s *safeLobby) init() {
	// 125 * 64 = 1kb
	s.process = make(chan func(floorMap), 125)
	s.registeredRooms = make(floorMap)
}

func (s *safeLobby) run() {
	for runner := range s.process {
		runner(s.registeredRooms)
	}
}

var roomLobby safeLobby

type Nroom struct {
	Name        string `json:"name"`
	MemberCount int    `json:"memberCount"`
}

func ActiveRooms(c *gin.Context) {

	result := make(chan []Nroom)
	defer close(result)

	arg := func(fm floorMap) {
		roomacc := make([]Nroom, 0, len(fm))
		for _, room := range fm {
			count := len(room.members)
			roomacc = append(roomacc, Nroom{Name: room.name, MemberCount: count})
		}
		result <- roomacc
	}

	roomLobby.process <- arg
	c.JSON(200, <-result)
}

func RoomSetup(c *gin.Context) {

	// extract queries and turn into regular request response function
	conn, err := upgrade(c.Writer, c.Request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var name string
	var room string
	if name, room = c.Query("userID"), c.Query("roomID"); name != "" && room != "" {
		success := make(chan int)

		arg := func(fm floorMap) {
			var chatroom *Room
			var ok bool
			if chatroom, ok = fm[room]; !ok {
				chatroom = new(Room)
				chatroom.OpenChatRoom(room)
				fm[room] = chatroom
			}
			user := new(Member)
			user.JoinRoom(name, chatroom, conn)
			success <- 200
		}

		roomLobby.process <- arg
		c.Status(<-success)

	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func init() {
	roomLobby.init()
	go roomLobby.run()
}
