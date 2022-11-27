package chatrooms

import (
	"encoding/json"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var testEngine *gin.Engine

func preflight() {
	Init()
	gin.SetMode(gin.TestMode)
	testEngine = gin.Default()
	testEngine.GET("/test", GinRoute)
}

func TestMain(m *testing.M) {
	preflight()
	code := m.Run()
	os.Exit(code)
}

func TestMessageSent(t *testing.T) {
	testServ := httptest.NewServer(testEngine)
	u := "ws" + strings.TrimPrefix(testServ.URL, "http")

	// 2 users join the room
	ws, _, err := websocket.DefaultDialer.Dial(u+"/test?userID=test1&roomID=test", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	defer ws.Close()
	ws2, _, err := websocket.DefaultDialer.Dial(u+"/test?userID=test2&roomID=test", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer ws2.Close()

	// send message to the connected room
	originalMessage := makeMessage("test1", "hello all")
	byteMessage, err := originalMessage.Marshal()
	if err != nil {
		t.Fatal(err.Error())
	}
	if err = ws.WriteMessage(websocket.TextMessage, byteMessage); err != nil {
		t.Fatal(err.Error())
	}

	// recieve and decode message
	_, recievedMessage, err := ws2.ReadMessage()
	if err != nil {
		t.Fatal(err.Error())
	}

	decodedMessage := new(Message)
	err = json.Unmarshal(recievedMessage, decodedMessage)
	if err != nil {
		t.Fatal(err.Error())
	}

	// do the messages match?
	if !reflect.DeepEqual(*decodedMessage, originalMessage) {
		t.Fatal(decodedMessage)
	}
}
