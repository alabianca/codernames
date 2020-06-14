package models

type CardType int

const RED = CardType(1)
const BLUE = CardType(4)
const ASSASSIN = CardType(0)
const CIVILIAN = CardType(7)

type Card struct {
	ID       string   `json:"id,omitempty"`
	X1       int32    `json:"x1"`
	Y1       int32    `json:"y1"`
	X2       int32    `json:"x2"`
	Y2       int32    `json:"y2"`
	Content  string   `json:"content"`
	CardType CardType `json:"type"`
	Active   bool     `json:"active"`
}

type Game struct {
	ID    string `json:"_id,omitempty"`
	Cards []Card `json:"cards"`
}
