package onu

import "gonu-server/onu/cards"

type CardPlacedEvent struct {
	BaseEvent
	Card cards.Card `json:"card"`
}

func NewCardPlacedEvent(card cards.Card) *CardPlacedEvent {
	return &CardPlacedEvent{
		BaseEvent: BaseEvent{Name: "CardPlacedEvent"},
		Card:      card,
	}
}
