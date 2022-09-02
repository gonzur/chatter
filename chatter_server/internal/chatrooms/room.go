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
	writeWait = 10 * time.Second
)

type Message struct {
	Sender string
	Data   []byte
}

func makeMessage(id string, value string) Message {
	message := Message{}

	message.Sender = id
	message.Data = []byte(value)

	return message
}

func (m Message) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

type Member struct {
	ID   string
	room *Room
	conn *websocket.Conn
	send chan Message
}

func (m *Member) Init(room *Room, conn *websocket.Conn) {
	m.conn = conn
	m.send = make(chan Message)
	m.room = room
	m.room.join <- m
	go m.OpenReciever()
	go m.OpenSender()
}

func (m *Member) OpenReciever() {
	defer func() {
		m.conn.Close()
		m.room.leave <- m
	}()

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

		readyMessage := makeMessage(m.ID, string(message))
		m.room.cast <- readyMessage
	}
}

func (m *Member) OpenSender() {
	pingTick := time.NewTicker(pingTime)
	defer func() {
		pingTick.Stop()
		m.conn.Close()
	}()
	for {
		select {
		case message, ok := <-m.send:
			if m.conn.SetWriteDeadline(time.Now().Add(writeWait)) != nil {
				return
			}
			if !ok {
				if m.conn.WriteMessage(websocket.CloseMessage, nil) != nil {
					return
				}
				return
			}
			if m.conn.WriteMessage(websocket.TextMessage, message.Data) != nil {
				return
			}
		case <-pingTick.C:
			if m.conn.SetWriteDeadline(time.Now().Add(writeWait)) != nil {
				return
			}
			if m.conn.WriteMessage(websocket.PingMessage, nil) != nil {
				return
			}
		}
	}

}

type Room struct {
	// should eventually hold a set of permissions not bool
	members map[*Member]bool
	join    chan *Member
	leave   chan *Member
	cast    chan Message
}

func (r *Room) Init() {
	r.members = make(map[*Member]bool)
	r.join = make(chan *Member)
	r.leave = make(chan *Member)
	r.cast = make(chan Message)
}

func (r *Room) Run() {
	for {
		select {
		case member := <-r.join:
			r.members[member] = true

		case member := <-r.leave:
			goodbye := makeMessage("", "Goodbye")
			r.cast <- goodbye
			delete(r.members, member)

		case message := <-r.cast:
			for m := range r.members {
				if m.ID != message.Sender {
					m.send <- message
				}

			}
		}
	}
}
