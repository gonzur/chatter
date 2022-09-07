package chatrooms

import (
	"errors"

	"github.com/gorilla/websocket"
)

type rooms struct {
	registeredRooms map[string]*Room
}

func makeEmptyRooms() rooms {
	return rooms{registeredRooms: make(map[string]*Room)}

}

var activeRooms rooms

func Create(name string) error {
	if _, ok := activeRooms.registeredRooms[name]; ok {
		return errors.New("exists")
	}
	room := new(Room)
	room.Init(name)
	activeRooms.registeredRooms[name] = room
	return nil

}

func Join(name string, conn *websocket.Conn) error {
	if room, ok := activeRooms.registeredRooms[name]; ok {
		mem := new(Member)
		mem.Init(room, conn)
		return nil
	}
	return errors.New("does not exist")

}

// TODO: upgrade passed connection to websocket connection and join room
func Upgrade() {

}

func Init() {
	activeRooms = makeEmptyRooms()
}
