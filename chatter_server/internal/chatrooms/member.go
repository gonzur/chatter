package chatrooms

import (
	"chatter-server/internal/chatrooms/sockconst"
	"time"

	"github.com/gorilla/websocket"
)

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

	m.conn.SetReadLimit(sockconst.ReadLimit)
	if m.conn.SetReadDeadline(time.Now().Add(sockconst.PongTime)) != nil {
		return
	}

	m.conn.SetPongHandler(func(string) error {
		return m.conn.SetReadDeadline(time.Now().Add(sockconst.PongTime))

	})

	for {
		messageType, message, err := m.conn.ReadMessage()
		if messageType == websocket.CloseMessage {
			return
		}
		if err != nil {
			return
		}

		readyMessage := makeMessage(m.ID, string(message))
		m.room.cast <- readyMessage
	}
}

func (m *Member) OpenSender() {
	pingTick := time.NewTicker(sockconst.PingTime)
	defer func() {
		pingTick.Stop()
		m.conn.Close()
	}()
	for {
		select {
		case message, ok := <-m.send:
			if m.conn.SetWriteDeadline(time.Now().Add(sockconst.WriteWait)) != nil {
				return
			}
			// TODO: log errors
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
			if m.conn.SetWriteDeadline(time.Now().Add(sockconst.WriteWait)) != nil {
				return
			}
			if m.conn.WriteMessage(websocket.PingMessage, nil) != nil {
				return
			}
		}
	}

}
