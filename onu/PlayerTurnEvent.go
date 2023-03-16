package onu

type PlayerTurnEvent struct {
	BaseEvent
	UUID string `json:"uuid"`
}

func NewPlayerTurnEvent(uuid string) *PlayerTurnEvent {
	return &PlayerTurnEvent{
		BaseEvent: BaseEvent{Name: "PlayerTurnEvent"},
		UUID:      uuid,
	}
}
