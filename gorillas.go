package gorillas

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

// Connection is just an alias to github.com/gorilla/websocket.Conn
type Connection *websocket.Conn

// Topic is an usual topic from messaging systems vocabulary
type Topic string

type connectionsMap map[Connection]bool

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
func (hub Gorillas) GetAllConnections() []Connection {
	return connectionsMapToSlice(hub.connectionsMap)
}

// AddConnection stores added connection in the connections hub
func (hub Gorillas) AddConnection(connection Connection) {
	hub.connectionsMap[connection] = true
}

// RemoveConnection removes specified connection from connections hub. Unsubscribes this connection from all topics as well.
func (hub Gorillas) RemoveConnection(connection Connection) {
	delete(hub.connectionsMap, connection)
	// @TODO Unsubscribe connection from
}

// Subscribe subscribes connection to the specified topic
func (hub Gorillas) Subscribe(connection Connection, topic Topic) {
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
func (hub Gorillas) Unsubscribe(connection Connection, topic Topic) {
	if _, topicExists := hub.subscriptionsMap[topic]; !topicExists {
		log.Fatal(fmt.Sprintf("Can not unsubscribe from unknown Topic [%s]", topic))
	}
	delete(hub.subscriptionsMap[topic], connection)
}

// GetSubscribedConnections returns all connections which are subscribed to the topic
func (hub Gorillas) GetSubscribedConnections(topic Topic) []Connection {
	connectionsMap, topicExists := hub.subscriptionsMap[topic]
	if !topicExists {
		log.Fatal(fmt.Sprintf("Can not get subscribers for every Topic from unknown Topic [%s]", topic))
	}
	return connectionsMapToSlice(connectionsMap)
}

// SendJSON allow send JSON to all connections subscribed to the topic
func (hub Gorillas) SendJSON(topic Topic, json interface{}) {
	subscribers := hub.GetSubscribedConnections(topic)
	for _, subscriber := range subscribers {
		connection := (*websocket.Conn)(subscriber)
		connection.WriteJSON(json)
	}
}

func connectionsMapToSlice(connectionsMap connectionsMap) []Connection {
	connections := make([]Connection, 0, len(connectionsMap))
	for k := range connectionsMap {
		connections = append(connections, k)
	}
	return connections
}
