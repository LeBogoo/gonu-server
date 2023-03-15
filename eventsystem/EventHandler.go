package eventsystem

import (
	"encoding/json"
	"fmt"
	"gonu-server/eventsystem/events"
	"reflect"

	"github.com/gorilla/websocket"
)

type EventHandler struct {
	eventTypes map[string]reflect.Type
	callbacks  map[string][]reflect.Value
}

func NewEventHandler() *EventHandler {
	return &EventHandler{
		eventTypes: make(map[string]reflect.Type),
		callbacks:  make(map[string][]reflect.Value),
	}
}

func (h *EventHandler) RegisterEvent(eventType interface{}) {
	name := reflect.TypeOf(eventType).Elem().Name()
	h.eventTypes[name] = reflect.TypeOf(eventType).Elem()
}

func (h *EventHandler) RegisterCallback(name string, callback interface{}) {
	h.callbacks[name] = append(h.callbacks[name], reflect.ValueOf(callback))
}

func (h *EventHandler) HandleMessage(message []byte, conn *websocket.Conn) error {
	// Parse the message into a Message struct
	var msg events.BaseEvent
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return err
	}

	// Check if the event type is registered
	eventType, ok := h.eventTypes[msg.Name]
	if !ok {
		return fmt.Errorf("unknown event type: %s", msg.Name)
	}

	// Create a new instance of the event type
	event := reflect.New(eventType).Interface()

	// Unmarshal the event data into the event instance
	err = json.Unmarshal(message, event)
	if err != nil {
		return err
	}

	// Call the callback functions for the event type
	for _, callback := range h.callbacks[msg.Name] {
		callback.Call([]reflect.Value{reflect.ValueOf(event), reflect.ValueOf(conn)})
	}

	return nil
}
