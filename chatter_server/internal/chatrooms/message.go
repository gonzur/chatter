package chatrooms

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
	SentOn  string `json:"sentOn"`
}

func convert12Hour(unFormTime time.Time) string {
	hour := unFormTime.Hour()
	if hour > 12 {
		hour = hour - 12
	}

	if hour == 0 {
		hour = 12;
	}

	meridian := "am"
	if unFormTime.Hour() >= 12 {
		meridian = "pm"
	}

	return fmt.Sprintf("%d:%02d %s", hour, unFormTime.Minute(), meridian)
}

func makeMessage(senderID string, text string) Message {
	message := Message{}

	message.Sender = senderID
	message.Message = text
	message.SentOn = convert12Hour(time.Now())

	return message
}

func (m Message) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
