package chatrooms

import "time"

const (
	readLimit = 512
	pingTime  = 50 * time.Second
	pongTime  = 60 * time.Second
	writeWait = 10 * time.Second
)
