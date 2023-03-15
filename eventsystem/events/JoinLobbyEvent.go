package events

type JoinLobbyEvent struct {
	BaseEvent
	LobbyCode string `json:"lobbyCode"`
	Username  string `json:"username"`
}

func NewJoinLobbyEvent(lobbyCode string, username string) *JoinLobbyEvent {
	return &JoinLobbyEvent{
		BaseEvent: BaseEvent{Name: "JoinLobbyEvent"},
		LobbyCode: lobbyCode,
		Username:  username,
	}
}
