package onu

type PlayerlistPlayer struct {
	Username   string `json:"username"`
	Uuid       string `json:"uuid"`
	CardCount  int    `json:"cardCount"`
	Active     bool   `json:"active"`
	Spectating bool   `json:"spectating"`
}

type UpdatePlayerlistEvent struct {
	BaseEvent
	Playerlist []*PlayerlistPlayer `json:"playerlist"`
}

func NewUpdatePlayerlistEvent(playerlist []*Player) *UpdatePlayerlistEvent {
	customPlayerlist := make([]*PlayerlistPlayer, len(playerlist))
	for i, player := range playerlist {
		customPlayerlist[i] = &PlayerlistPlayer{
			Username:   player.Username,
			Uuid:       player.UserId,
			CardCount:  len(player.Cards),
			Active:     player.Game.GetActivePlayer() == player,
			Spectating: player.Spectating,
		}
	}

	return &UpdatePlayerlistEvent{
		BaseEvent:  BaseEvent{Name: "UpdatePlayerlistEvent"},
		Playerlist: customPlayerlist,
	}
}
