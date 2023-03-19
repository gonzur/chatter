package chatrooms

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Member struct {
	ID   string
	room *Room
	conn *websocket.Conn
	send chan Message
}

func (m *Member) JoinRoom(userID string, room *Room, conn *websocket.Conn) {
	m.conn = conn
	m.send = make(chan Message)
	m.room = room
	m.ID = userID
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
		readyMessage := &Message{}

		err := m.conn.ReadJSON(readyMessage)

		if err != nil {
			log.Println(err.Error())
			return
		}

		m.room.cast <- *readyMessage
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

			// TODO: log errors
			if !ok {
				if m.conn.WriteMessage(websocket.CloseMessage, nil) != nil {
					return
				}
				return
			}

			if err := m.conn.WriteJSON(message); err != nil {
				log.Println(err.Error())
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
