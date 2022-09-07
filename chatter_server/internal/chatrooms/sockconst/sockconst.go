package sockconst

import "time"

const (
	ReadLimit = 512
	PingTime  = 60 * time.Second
	PongTime  = 50 * time.Second
	WriteWait = 10 * time.Second
)
