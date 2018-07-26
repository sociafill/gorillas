package gorillas

import (
	"fmt"
	"log"
)

// ConnectionInterface is an interface which MUST be implemented by WebSockets library.
// As example: github.com/gorilla/websocket.Conn
type ConnectionInterface interface {
	WriteJSON(interface{}) error
}

// Topic is an usual topic from messaging systems vocabulary
type Topic string

type connectionsMap map[ConnectionInterface]bool

// Gorillas is main structure for the connections management
type Gorillas struct {
	connectionsMap   connectionsMap
	subscriptionsMap map[Topic]connectionsMap
}

// NewGorillas creates instance of Gorillas
func NewGorillas() Gorillas {
	hub := Gorillas{}
	hub.connectionsMap = make(connectionsMap)
	hub.subscriptionsMap = make(map[Topic]connectionsMap)
	return hub
}

// GetAllConnections returns all registered connections from Gorillas hub
func (hub Gorillas) GetAllConnections() []ConnectionInterface {
	return connectionsMapToSlice(hub.connectionsMap)
}

// AddConnection stores added connection in the connections hub
func (hub Gorillas) AddConnection(connection ConnectionInterface) {
	hub.connectionsMap[connection] = true
}

// RemoveConnection removes specified connection from connections hub. Unsubscribes this connection from all topics as well.
func (hub Gorillas) RemoveConnection(connection ConnectionInterface) {
	delete(hub.connectionsMap, connection)
	// @TODO Unsubscribe connection from
}

// Subscribe subscribes connection to the specified topic
func (hub Gorillas) Subscribe(connection ConnectionInterface, topic Topic) {
	if _, connExists := hub.connectionsMap[connection]; !connExists {
		log.Fatal("Can not subscribe unknown connection to the Topic")
		return
	}
	if _, topicExists := hub.subscriptionsMap[topic]; !topicExists {
		hub.subscriptionsMap[topic] = make(connectionsMap)
	}
	hub.subscriptionsMap[topic][connection] = true
}

// Unsubscribe unsubscribes connection from the specified topic
func (hub Gorillas) Unsubscribe(connection ConnectionInterface, topic Topic) {
	if _, topicExists := hub.subscriptionsMap[topic]; !topicExists {
		log.Fatal(fmt.Sprintf("Can not unsubscribe from unknown Topic [%s]", topic))
	}
	delete(hub.subscriptionsMap[topic], connection)
}

// GetSubscribedConnections returns all connections which are subscribed to the topic
func (hub Gorillas) GetSubscribedConnections(topic Topic) []ConnectionInterface {
	connectionsMap, topicExists := hub.subscriptionsMap[topic]
	if !topicExists {
		// @TODO Add adequate handling for case when nobody is subscribed to topic
		log.Fatal(fmt.Sprintf("Can not get subscribers for every Topic from unknown Topic [%s]", topic))
	}
	return connectionsMapToSlice(connectionsMap)
}

// SendJSON allow send JSON to all connections subscribed to the topic
func (hub Gorillas) SendJSON(topic Topic, json interface{}) {
	subscribers := hub.GetSubscribedConnections(topic)
	for _, subscriber := range subscribers {
		subscriber.WriteJSON(json)
	}
}

func connectionsMapToSlice(connectionsMap connectionsMap) []ConnectionInterface {
	connections := make([]ConnectionInterface, 0, len(connectionsMap))
	for k := range connectionsMap {
		connections = append(connections, k)
	}
	return connections
}
