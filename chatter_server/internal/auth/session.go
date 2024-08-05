package auth

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
)

type Session struct {
	Key      []byte
	Username string
}

var userSessions sync.Map

func AddSession(username string) []byte {
	sess := new(Session)

	key := newSessionKey()

	sess.Username = username
	sess.Key = newSessionKey()

	userSessions.Store(key, sess)
	return key
}

func FetchSession(key []byte) *Session {
	if sess, ok := userSessions.Load(key); ok {
		return sess.(*Session)
	}
	return nil
}

func newSessionKey() []byte {
	randStr := make([]byte, 196)
	returnStr := make([]byte, 196)
	rand.Reader.Read(randStr)
	base64.StdEncoding.Encode(returnStr, randStr)

	return returnStr
}
