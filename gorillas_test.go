package gorillas

import (
	"testing"

	"github.com/gorilla/websocket"
)

func TestConstructor(t *testing.T) {
	NewGorillas()
}

func TestAddConnectionSuccessfully(t *testing.T) {
	gorillas := NewGorillas()
	connection := websocket.Conn{}
	gorillas.AddConnection(&connection)
}

func TestMessagesDeliveryToSubscribers(t *testing.T) {

	topic := Topic("test-topic")
	gorillas := NewGorillas()
	connection := websocket.Conn{}
	gorillas.AddConnection(&connection)
	gorillas.Subscribe(&connection, topic)
	data := "Hello, world"
	gorillas.SendJSON(topic, data)
}
