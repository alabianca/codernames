package core

import (
	"bufio"
	"errors"
	"io"
)

const GameSize = 25

type GameDAL interface {
	Activate(id string, content string, doc interface{}) error
	Create(interface{})  error
	Get(interface{}) error
}

type Coordinate struct {
	X1 int32
	Y1 int32
	X2 int32
	Y2 int32
}

type Card interface {
	GetContent() string
	GetType() int
	IsActive() bool
	Coords() Coordinate
}

func ProcessWords(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)

	words := make([]string, 0, GameSize)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}


	if len(words) < GameSize {
		return nil, errors.New("invalid size. need 25 words")
	}

	return words, nil
}
