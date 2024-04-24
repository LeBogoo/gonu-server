package cards

import (
	"math/rand"
)

type Preset struct {
	Colors []Color
	Types  []string
}

func (c *Preset) RandomCard() *Card {
	color := c.Colors[rand.Intn(len(c.Colors))]
	cardType := c.Types[rand.Intn(len(c.Types))]

	return NewCard(cardType, color)
}
