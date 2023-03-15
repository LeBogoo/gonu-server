package onu

import (
	"fmt"
	"gonu-server/eventsystem/events"
)

type Game struct {
	LobbyCode string
	Players   []*Player
	Admin     *Player
	Settings  map[string]events.OnuSetting
}

func NewGame(lobbyCode string) *Game {

	game := &Game{
		LobbyCode: lobbyCode,
		Settings:  make(map[string]events.OnuSetting),
	}

	game.SetSetting(events.OnuSetting{
		Name:     "Card amount",
		Value:    "7",
		Defaults: []string{"5", "7", "10", "15", "20"},
	})

	game.SetSetting(events.OnuSetting{
		Name:     "Gamemode",
		Value:    "Classic",
		Defaults: []string{"Lite", "Classic", "Special"},
	})

	game.SetSetting(events.OnuSetting{
		Name:     "Max Players",
		Value:    "10",
		Defaults: []string{"10", "20", "30", "40"},
	})

	return game
}

func (g *Game) AddPlayer(player *Player) {
	g.Players = append(g.Players, player)
}

func (g *Game) SetAdmin(player *Player) {
	g.Admin = player

	fmt.Printf("New admin for game %s is %s\n", g.LobbyCode, g.Players[0].UserId)
	for _, p := range g.Players {
		p.Ws.WriteJSON(events.NewUpdateAdminEvent(player.UserId))
	}
}

func (g *Game) RemovePlayer(player *Player) {
	for i, p := range g.Players {
		if p == player {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			break
		}
	}

	if g.Admin == player {
		// the new admin should be the first player in the list. if the list is empty, the game should be deleted
		if len(g.Players) > 0 {
			fmt.Println("Old admin left.")
			g.SetAdmin(g.Players[0])
		} else {
			fmt.Println("Deleting game", g.LobbyCode)
			games := player.Games
			delete(*games, g.LobbyCode)
		}
	}

}

func (g *Game) GetPlayerById(userId string) *Player {
	for _, p := range g.Players {
		if p.UserId == userId {
			return p
		}
	}

	return nil
}

func (g *Game) BroadcastEvent(event interface{}) {
	for _, p := range g.Players {
		p.Ws.WriteJSON(event)
	}
}

func (g *Game) BroadcastSettings() {
	settingsEvent := events.NewSettingsChangedEvent(g.Settings)
	g.BroadcastEvent(settingsEvent)
}

func (g *Game) SetSetting(setting events.OnuSetting) {
	g.Settings[setting.Name] = setting
}
