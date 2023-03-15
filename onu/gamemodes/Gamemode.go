package gamemodes

import "gonu-server/onu/cards"

type Gamemode interface {
	GetName() string
	GetDescription() string
	RandomCard() cards.Card
	CompareCards(cards.Card, cards.Card) bool
}
