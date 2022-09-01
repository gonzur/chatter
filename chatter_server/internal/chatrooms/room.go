package chatrooms

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

const (
	readLimit = 512
	pingTime  = 60 * time.Second
	pongTime  = (pingTime * 9) / 10
)

func makeMessage(id string, value string) ([]byte, error) {
	message := map[string][]byte{
		"id":      []byte(id),
		"message": []byte(value),
	}
	return json.Marshal(message)
}

type Member struct {
	ID      string
	room    *Room
	conn    *websocket.Conn
	Send    chan []byte
	Recieve chan []byte
}

func (m *Member) Init(conn *websocket.Conn) {
	m.conn = conn
	m.Send = make(chan []byte)
	m.Recieve = make(chan []byte)
}

func (m Member) OpenReciever() {
	m.conn.SetReadLimit(readLimit)
	if m.conn.SetReadDeadline(time.Now().Add(pongTime)) != nil {
		return
	}

	m.conn.SetPongHandler(func(string) error {
		return m.conn.SetReadDeadline(time.Now().Add(pongTime))

	})

	for {
		_, message, err := m.conn.ReadMessage()
		if err != nil {
			return
		}

		readyMessage, err := makeMessage(m.ID, string(message))
		if err != nil {
			return
		}

		m.room.cast <- readyMessage
	}
}

type Room struct {
	// should eventually hold a set of permissions not bool
	members map[*Member]bool
	join    chan *Member
	leave   chan *Member
	cast    chan []byte
}

func (r *Room) Init() {
	r.members = make(map[*Member]bool)
	r.join = make(chan *Member)
	r.leave = make(chan *Member)
	r.cast = make(chan []byte)
}

func (r *Room) Run() {
	for {
		select {
		case member := <-r.join:
			r.members[member] = true

		case member := <-r.leave:

			goodbye, err := makeMessage("", "Goodbye")
			if err != nil {
				continue
			}
			r.cast <- goodbye
			delete(r.members, member)

		case data := <-r.cast:
			dataMap := make(map[string][]byte)
			if json.Unmarshal(data, &dataMap) != nil {
				continue
			}

			id := string(dataMap["id"])
			for m := range r.members {
				if m.ID == id {
					continue
				}
				m.Send <- dataMap["data"]
			}
		}
	}
}
