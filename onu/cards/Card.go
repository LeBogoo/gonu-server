package cards

import (
	"github.com/google/uuid"
)

type Card struct {
	Type  string `json:"type"`
	Color Color  `json:"color"`
	Id    string `json:"id"`
}

func NewCard(cardType string, color Color) *Card {
	return &Card{
		Type:  cardType,
		Color: color,
		Id:    uuid.New().String(),
	}
}
