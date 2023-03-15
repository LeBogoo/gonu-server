package onu

import (
	"gonu-server/eventsystem"
	"gonu-server/eventsystem/events"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Player struct {
	Ws           *websocket.Conn
	EventHandler *eventsystem.EventHandler
}

func registerCallbacks(handler *eventsystem.EventHandler) {
	handler.RegisterEvent(&events.JoinLobbyEvent{})
	handler.RegisterEvent(&events.SettingsChangedEvent{})

	handler.RegisterCallback("JoinLobbyEvent", func(event *events.JoinLobbyEvent, conn *websocket.Conn) {
		log.Println("JoinLobbyEvent", event.Name, event.LobbyCode, event.Username)

		// generate a random UUID
		userId, _ := uuid.NewRandom()

		conn.WriteJSON(events.NewJoinedLobbyEvent(userId.String()))

		settings := make(map[string]events.OnuSettings)
		settings["CardAmount"] = events.OnuSettings{
			Name:     "Card amount",
			Value:    "7",
			Defaults: []string{"5", "7", "10", "15", "20"},
		}

		settings["GameMode"] = events.OnuSettings{
			Name:     "Gamemode",
			Value:    "Classic",
			Defaults: []string{"Classic", "Lite", "Special"},
		}

		settings["DevMode"] = events.OnuSettings{
			Name:     "Devmode",
			Value:    "Normal",
			Defaults: []string{"Normal", "Verbose", "Debug"},
		}

		conn.WriteJSON(events.NewSettingsChangedEvent(settings))

		conn.WriteJSON(events.NewUpdateAdminEvent(userId.String()))
	})

	handler.RegisterCallback("SettingsChangedEvent", func(event *events.SettingsChangedEvent, conn *websocket.Conn) {
		log.Println("SettingsChangedEvent", event.Name, event.Settings)
	})
}

func NewConnection(ws *websocket.Conn) *Player {
	handler := eventsystem.NewEventHandler()

	registerCallbacks(handler)

	player := &Player{
		Ws:           ws,
		EventHandler: handler,
	}

	go func() {
		defer func() {
			ws.Close()
		}()

		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				break
			}

			handler.HandleMessage(msg, ws)
		}
	}()

	return player
}
