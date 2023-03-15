package events

import (
	"encoding/json"
)

type BaseEvent struct {
	Name string `json:"name"`
}

func (e *BaseEvent) Stringify() string {
	eventData, _ := json.Marshal(e)

	return string(eventData)
}

func ParseEvent(eventString string) (*BaseEvent, error) {
	eventData := &BaseEvent{}
	err := json.Unmarshal([]byte(eventString), eventData)
	if err != nil {
		return nil, err
	}
	return eventData, nil
}
