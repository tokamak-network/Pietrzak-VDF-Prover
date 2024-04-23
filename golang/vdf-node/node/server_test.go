package node

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"testing"
)

func TestWebSocketConnection(t *testing.T) {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8545", Path: "/ws"}

	log.Printf("Connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		t.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	defer c.Close()

	err = c.WriteMessage(websocket.TextMessage, []byte("hello"))
	if err != nil {
		t.Fatalf("Failed to send message to WebSocket: %v", err)
	}

	_, message, err := c.ReadMessage()
	if err != nil {
		t.Fatalf("Failed to read message from WebSocket: %v", err)
	}

	log.Printf("Received: %s", message)
}
