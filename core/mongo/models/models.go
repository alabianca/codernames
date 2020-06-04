package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CardType int

const RED = CardType(1)
const BLUE = CardType(4)
const ASSASSIN = CardType(0)
const CIVILIAN = CardType(7)

type Card struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	X1       int32              `bson:"x1"`
	Y1       int32              `bson:"y1"`
	X2       int32              `bson:"x2"`
	Y2       int32              `bson:"y2"`
	Content  string             `bson:"content"`
	CardType CardType           `bson:"type"`
	Active   bool               `bson:"active"`
}

type Game struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Cards []Card             `bson:"cards"`
}
