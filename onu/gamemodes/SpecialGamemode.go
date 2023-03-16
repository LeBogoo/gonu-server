package gamemodes

import "gonu-server/onu/cards"

type SpecialGamemode struct {
	Gamemode
}

func NewSpecialGamemode() *SpecialGamemode {
	return &SpecialGamemode{}
}

func (c *SpecialGamemode) GetName() string {
	return "Special"
}

func (c *SpecialGamemode) GetDescription() string {
	return "Special gamemode"
}

func (c *SpecialGamemode) RandomCard() cards.Card {
	return *cards.NewCard("0", *cards.ColorFrom("g"))
}

func (c *SpecialGamemode) CompareCards(cards.Card, cards.Card) bool {
	return true
}
