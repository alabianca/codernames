package client

import (
	"errors"
	"github.com/alabianca/codernames/core"
	"math/rand"
	"time"
)

type card struct {
	ctype    int
	content  string
	isActive bool
	x1       int32
	y1       int32
	x2       int32
	y2       int32
}

func (c *card) GetContent() string {
	return c.content
}

func (c *card) GetType() int {
	return c.ctype
}

func (c *card) IsActive() bool {
	return c.isActive
}

func (c *card) Coords() core.Coordinate {
	return core.Coordinate{
		X1: c.x1,
		X2: c.x2,
		Y1: c.y1,
		Y2: c.y2,
	}
}

func newCard(content string, ctype int) *card {
	return &card{
		content: content,
		ctype:   ctype,
	}
}

const red = 1
const blue = 4
const white = 7
const black = 0

func generate25Cards(words []string) ([]*card, error) {
	if len(words) < 25 {
		return nil, errors.New("25 cards are needed")
	}

	cards := []*card{
		newCard(words[0], 1),
		newCard(words[1], 1),
		newCard(words[2], 1),
		newCard(words[3], 1),
		newCard(words[4], 1),
		newCard(words[5], 1),
		newCard(words[6], 1),
		newCard(words[7], 1),
		newCard(words[8], 1),
		newCard(words[9], 4),
		newCard(words[10], 4),
		newCard(words[11], 4),
		newCard(words[12], 4),
		newCard(words[13], 4),
		newCard(words[14], 4),
		newCard(words[15], 4),
		newCard(words[16], 4),
		newCard(words[17], 7),
		newCard(words[18], 7),
		newCard(words[19], 7),
		newCard(words[20], 7),
		newCard(words[21], 7),
		newCard(words[22], 7),
		newCard(words[23], 7),
		newCard(words[24], 0),
	}

	shuffle(cards)

	return cards, nil

}

func shuffle(words []*card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
}
