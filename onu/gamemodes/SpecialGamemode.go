package gamemodes

import (
	"gonu-server/onu/cards"
	"math/rand"
)

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
	presets := []cards.Preset{
		*cards.NewClassicPreset(),
		*cards.NewWishPreset(),
	}

	preset := presets[rand.Intn(len(presets))]

	return *preset.RandomCard()
}

func (c *SpecialGamemode) CompareCards(cards.Card, cards.Card) bool {
	return true
}
