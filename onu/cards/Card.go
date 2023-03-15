package cards

import (
	"github.com/google/uuid"
)

type Card struct {
	Type  string
	Color Color
	Id    string
}

func NewCard(cardType string, color Color) *Card {
	return &Card{
		Type:  cardType,
		Color: color,
		Id:    uuid.New().String(),
	}
}
