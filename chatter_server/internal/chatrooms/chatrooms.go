package chatrooms

import (
	"errors"
	"net/http"

	"github.com/gorilla/websocket"
)

type rooms struct {
	registeredRooms map[string]*Room
}

func makeEmptyRooms() rooms {
	return rooms{registeredRooms: make(map[string]*Room)}

}

var activeRooms rooms

var socketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Create(name string) error {
	if _, ok := activeRooms.registeredRooms[name]; ok {
		return errors.New("exists")
	}
	room := new(Room)
	room.Init(name)
	activeRooms.registeredRooms[name] = room
	return nil

}

func Join(roomID string, userID string, conn *websocket.Conn) error {
	if room, ok := activeRooms.registeredRooms[roomID]; ok {
		mem := new(Member)
		mem.Init(userID, room, conn)
		return nil
	}
	return errors.New("does not exist")

}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return socketUpgrade.Upgrade(w, r, nil)
}

func Init() {
	activeRooms = makeEmptyRooms()
}
