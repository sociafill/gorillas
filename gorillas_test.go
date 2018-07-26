package gorillas

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sociafill/gorillas/mocks"
)

type DummyConnection struct {
	messages []interface{}
}

func (connection DummyConnection) WriteJSON(v interface{}) error {
	return nil
}

func TestConstructor(t *testing.T) {
	NewGorillas()
}

func TestAddConnectionSuccessfully(t *testing.T) {
	gorillas := NewGorillas()
	// var connection ConnectionInterface
	connection := &DummyConnection{}

	gorillas.AddConnection(connection)
}

func TestMessagesDeliveryToSubscribers(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockConnection := mocks.NewMockConnectionInterface(mockCtrl)
	testGorillas := NewGorillas()
	testGorillas.AddConnection(mockConnection)
	topic := Topic("test-topic")
	testGorillas.Subscribe(mockConnection, topic)
	data := "Hello, world"
	mockConnection.EXPECT().WriteJSON(data).Return(nil).Times(1)
	testGorillas.SendJSON(topic, data)
}
