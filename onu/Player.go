package onu

import (
	"fmt"
	"gonu-server/eventsystem"
	"gonu-server/eventsystem/events"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Player struct {
	Ws           *websocket.Conn
	UserId       string
	Username     string
	EventHandler *eventsystem.EventHandler
	Games        *map[string]*Game
	Game         *Game
}

func (p *Player) registerCallbacks(handler *eventsystem.EventHandler) {
	handler.RegisterEvent(&events.JoinLobbyEvent{})
	handler.RegisterEvent(&events.SettingsChangedEvent{})

	handler.RegisterCallback("JoinLobbyEvent", func(event *events.JoinLobbyEvent, conn *websocket.Conn) {
		p.Username = event.Username

		game := (*p.Games)[event.LobbyCode]
		created := game == nil

		if game == nil {
			game = NewGame(event.LobbyCode)
			(*p.Games)[event.LobbyCode] = game
		}

		game.AddPlayer(p)

		if created {
			go func() {
				time.Sleep(100 * time.Millisecond)
				game.SetAdmin(p)
			}()
		}

		p.Game = game

		fmt.Println(game.LobbyCode, game.Players)

		conn.WriteJSON(events.NewJoinedLobbyEvent(p.UserId))
		conn.WriteJSON(events.NewSettingsChangedEvent(game.Settings))
	})

	handler.RegisterCallback("SettingsChangedEvent", func(event *events.SettingsChangedEvent, conn *websocket.Conn) {
		if p.Game == nil || p.Game.Admin != p {
			return
		}

		for _, setting := range event.Settings {
			p.Game.SetSetting(setting)
		}

		p.Game.BroadcastSettings()
	})
}

func NewPlayer(ws *websocket.Conn, games *map[string]*Game) *Player {
	handler := eventsystem.NewEventHandler()

	player := &Player{
		Ws:           ws,
		EventHandler: handler,
		Games:        games,
		UserId:       uuid.New().String(),
	}

	player.registerCallbacks(handler)

	go func() {
		defer func() { ws.Close() }()

		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					if player.Game != nil {
						player.Game.RemovePlayer(player)
					}
				} else {
					log.Println("Error reading message:", err)
				}
				break
			}

			handler.HandleMessage(msg, ws)
		}
	}()

	return player
}
