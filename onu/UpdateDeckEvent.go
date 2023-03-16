package onu

import "gonu-server/onu/cards"

type UpdateDeckEvent struct {
	BaseEvent
	Deck []cards.Card `json:"deck"`
}

func NewUpdateDeckEvent(deck []cards.Card) *UpdateDeckEvent {
	return &UpdateDeckEvent{
		BaseEvent: BaseEvent{Name: "UpdateDeckEvent"},
		Deck:      deck,
	}
}
