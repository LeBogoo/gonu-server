package onu

type UpdatePlayerlistEvent struct {
	BaseEvent
	UUID string `json:"uuid"`
}

func NewUpdatePlayerlistEvent(playerlist []interface{}) *UpdatePlayerlistEvent {
	return &UpdatePlayerlistEvent{
		BaseEvent: BaseEvent{Name: "UpdatePlayerlistEvent"},
	}
}
