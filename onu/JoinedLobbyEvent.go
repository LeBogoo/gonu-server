package onu

type JoinedLobbyEvent struct {
	BaseEvent
	UUID string `json:"uuid"`
}

func NewJoinedLobbyEvent(uuid string) *JoinedLobbyEvent {
	return &JoinedLobbyEvent{
		BaseEvent: BaseEvent{Name: "JoinedLobbyEvent"},
		UUID:      uuid,
	}
}
