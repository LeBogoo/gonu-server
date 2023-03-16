package onu

type UpdatePlayerlistEvent struct {
	BaseEvent
	Playerlist []*Player `json:"playerlist"`
}

func NewUpdatePlayerlistEvent(playerlist []*Player) *UpdatePlayerlistEvent {
	return &UpdatePlayerlistEvent{
		BaseEvent:  BaseEvent{Name: "UpdatePlayerlistEvent"},
		Playerlist: playerlist,
	}
}
