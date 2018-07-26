package gorillas

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sociafill/gorillas/mocks"
)

func TestConstructor(t *testing.T) {
	NewGorillas()
}

func TestAddConnectionSuccessfully(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockConnection := mocks.NewMockConnectionInterface(mockCtrl)
	hub := NewGorillas()
	hub.AddConnection(mockConnection)
}

func TestMessagesDeliveryToSubscribers(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockConnection := mocks.NewMockConnectionInterface(mockCtrl)
	hub := NewGorillas()
	hub.AddConnection(mockConnection)
	topic := Topic("test-topic")
	hub.Subscribe(mockConnection, topic)
	data := "Hello, world"
	mockConnection.EXPECT().WriteJSON(data).Return(nil).Times(1)
	hub.SendJSON(topic, data)
}
