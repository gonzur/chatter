package chatrooms

import "encoding/json"

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
