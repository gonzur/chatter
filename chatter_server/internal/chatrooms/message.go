package chatrooms

import (
	"encoding/json"
	"time"
)

type Message struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
	SentOn  string `json:"sentOn"`
}

func makeMessage(id string, value string) Message {
	message := Message{}

	message.Sender = id
	message.Message = value
	message.SentOn = time.Now().Format(time.RFC1123)

	return message
}

func (m Message) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
