package chatrooms

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type rooms struct {
	registeredRooms map[string]*Room
}

func makeEmptyRooms() rooms {
	return rooms{registeredRooms: make(map[string]*Room)}

}

var activeRooms rooms = makeEmptyRooms()

var socketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func createIfNotExist(name string) {
	if _, ok := activeRooms.registeredRooms[name]; ok {
		return
	}
	room := new(Room)
	room.OpenChatRoom(name)
	activeRooms.registeredRooms[name] = room
}

func Join(roomID string, userID string, conn *websocket.Conn) error {
	if room, ok := activeRooms.registeredRooms[roomID]; ok {
		mem := new(Member)
		mem.JoinRoom(userID, room, conn)
		return nil
	}
	return errors.New("does not exist")

}

func GinRoute(c *gin.Context) {

	// extract queries and turn into regular request response function
	conn, err := upgrade(c.Writer, c.Request)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var name string
	var room string
	if name, room = c.Query("userID"), c.Query("roomID"); name != "" && room != "" {

		// error here
		createIfNotExist(room)

		if err = Join(room, name, conn); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return socketUpgrade.Upgrade(w, r, nil)
}