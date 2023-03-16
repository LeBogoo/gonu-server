package onu

import (
	"fmt"
	"gonu-server/onu/cards"
	"gonu-server/onu/gamemodes"
	"strconv"
	"time"
)

type Game struct {
	LobbyCode    string
	Players      []*Player
	ActivePlayer int
	Admin        *Player
	GameMode     gamemodes.Gamemode
	Settings     map[string]OnuSetting
}

var modes = []gamemodes.Gamemode{
	gamemodes.NewClassicGamemode(),
	gamemodes.NewLiteGamemode(),
	gamemodes.NewSpecialGamemode(),
}

func NewGame(lobbyCode string) *Game {

	game := &Game{
		LobbyCode:    lobbyCode,
		ActivePlayer: 0,
		Settings:     make(map[string]OnuSetting),
	}

	game.SetSetting(OnuSetting{
		Name:     "Card amount",
		Value:    "7",
		Defaults: []string{"5", "7", "10", "15", "20"},
	})

	gamemodeNames := make([]string, len(modes))
	for i, gamemode := range modes {
		gamemodeNames[i] = gamemode.GetName()
	}

	game.SetSetting(OnuSetting{
		Name:     "Gamemode",
		Value:    gamemodeNames[0],
		Defaults: gamemodeNames,
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
		p.Ws.WriteJSON(NewUpdateAdminEvent(player.UserId))
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

func (g *Game) GetPlayers() ([]*Player, []*Player) {
	players := make([]*Player, 0)
	spectators := make([]*Player, 0)
	for _, p := range g.Players {
		if !p.Spectating {
			players = append(players, p)
		} else {
			spectators = append(spectators, p)
		}
	}
	return players, spectators
}

func (g *Game) GetActivePlayer() *Player {
	players, _ := g.GetPlayers()
	if len(players) != 0 {
		return players[g.ActivePlayer]
	}
	return nil
}

func (g *Game) BroadcastEvent(event interface{}) {
	for _, p := range g.Players {
		p.Ws.WriteJSON(event)
	}
}

func (g *Game) BroadcastSettings() {
	settingsEvent := NewSettingsChangedEvent(g.Settings)
	g.BroadcastEvent(settingsEvent)
}

func (g *Game) BroadcastPlayerlist() {
	playerlistEvent := NewUpdatePlayerlistEvent(g.Players)
	g.BroadcastEvent(playerlistEvent)
}

func (g *Game) SetSetting(setting OnuSetting) {
	g.Settings[setting.Name] = setting

	if setting.Name == "Gamemode" {
		g.SetGamemode(setting.Value)
	}
}

func (g *Game) SetGamemode(gamemode string) {
	g.GameMode = modes[0]
	for _, mode := range modes {
		if mode.GetName() == gamemode {
			g.GameMode = mode
			break
		}
	}

	fmt.Printf("Switched gamemode to %s\n", g.GameMode.GetDescription())
}

func (g *Game) Start() {
	fmt.Println("Starting game", g.LobbyCode)
	gameStartEvent := NewGameStartEvent()
	g.BroadcastEvent(gameStartEvent)

	g.ActivePlayer = 0
	for _, p := range g.Players {
		p.Spectating = false

		// generate cards
		p.Cards = make([]cards.Card, 0)
		cardAmount, _ := strconv.Atoi(g.Settings["Card amount"].Value)

		for i := 0; i < cardAmount; i++ {
			p.Cards = append(p.Cards, g.GameMode.RandomCard())
		}

		p.Ws.WriteJSON(NewUpdateDeckEvent(p.Cards))
	}

	g.BroadcastEvent(NewCardPlacedEvent(*cards.NewCard("w", *cards.ColorFrom("r"))))
	go func() {
		time.Sleep(1 * time.Second)
		g.BroadcastEvent(NewPlayerTurnEvent(g.GetActivePlayer().UserId))

	}()

	g.BroadcastPlayerlist()
}
