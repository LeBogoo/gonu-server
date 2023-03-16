package onu

type GameStartEvent struct {
	BaseEvent
}

func NewGameStartEvent() *GameStartEvent {
	return &GameStartEvent{
		BaseEvent: BaseEvent{Name: "GameStartEvent"},
	}
}
