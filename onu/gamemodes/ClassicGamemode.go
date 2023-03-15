package gamemodes

import "gonu-server/onu/cards"

type ClassicGamemode struct {
	Gamemode
}

func NewClassicGamemode() *ClassicGamemode {
	return &ClassicGamemode{}
}

func (c *ClassicGamemode) GetName() string {
	return "Classic"
}

func (c *ClassicGamemode) GetDescription() string {
	return "Classic gamemode"
}

func (c *ClassicGamemode) RandomCard() cards.Card {
	return *cards.NewCard("0", *cards.ColorFrom("r"))
}

func (c *ClassicGamemode) CompareCards(cards.Card, cards.Card) bool {
	return true
}
