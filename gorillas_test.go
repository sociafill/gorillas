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
