package gamemodes

import "gonu-server/onu/cards"

type LiteGamemode struct {
	Gamemode
}

func NewLiteGamemode() *LiteGamemode {
	return &LiteGamemode{}
}

func (c *LiteGamemode) GetName() string {
	return "Lite"
}

func (c *LiteGamemode) GetDescription() string {
	return "Lite gamemode"
}

func (c *LiteGamemode) RandomCard() cards.Card {
	return *cards.NewCard("0", *cards.ColorFrom("b"))
}

func (c *LiteGamemode) CompareCards(cards.Card, cards.Card) bool {
	return true
}
