package chatrooms

import "time"

const (
	readLimit = 512
	pingTime  = 60 * time.Second
	pongTime  = 50 * time.Second
	writeWait = 10 * time.Second
)
