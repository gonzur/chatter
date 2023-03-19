package chatrooms

import (
	"fmt"
	"log"
)

type Room struct {
	name string
	// should eventually hold a set of permissions not bool
	members map[*Member]bool
	join    chan *Member
	leave   chan *Member
	cast    chan Message
}

func (r *Room) OpenChatRoom(roomName string) {
	r.name = roomName
	r.members = make(map[*Member]bool)
	r.join = make(chan *Member)
	r.leave = make(chan *Member)
	r.cast = make(chan Message)
	go r.Run()
}

func (r *Room) transmitExclusive(message Message) {
	for m := range r.members {
		if m.ID != message.Sender {
			m.send <- message
		}
	}
}

func (r *Room) Run() {
	defer func() {
		err := recover()
		log.Println(err)
	}()

	for {
		select {
		case messenger := <-r.join:
			r.members[messenger] = true

		case messenger := <-r.leave:
			leaveMessage := fmt.Sprintf("%s left...", messenger.ID)
			goodbye := makeMessage("System", leaveMessage)
			r.transmitExclusive(goodbye)
			delete(r.members, messenger)

		case message := <-r.cast:
			r.transmitExclusive(message)
		}
	}
}
