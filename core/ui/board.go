package ui

import (
	"fmt"
	"github.com/alabianca/codernames/core"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type CardType ui.Color

const Russia = CardType(ui.ColorRed)
const USA = CardType(ui.ColorBlue)
const Civilian = CardType(ui.ColorYellow)
const Assassin = CardType(ui.ColorWhite)

type Card struct {
	paragraph *widgets.Paragraph
	CardType  CardType
	isActive  bool
}

func (c Card) GetContent() string {
	return c.paragraph.Text
}

func (c Card) GetType() int {
	return int(c.CardType)
}

func (c Card) IsActive() bool {
	return c.isActive
}

func (c Card) Coords() core.Coordinate {
	return core.Coordinate{}
}

func NewCard(text string, ct CardType, isActive bool) *Card {
	c := &Card{
		paragraph: widgets.NewParagraph(),
		CardType:  ct,
	}

	c.paragraph.Text = text
	if isActive {
		c.paragraph.BorderStyle = ui.Style{Fg: ui.Color(ct)}
		c.isActive = true
	} else {
		c.paragraph.BorderStyle = ui.Style{Fg: ui.Color(Assassin)}
	}

	return c

}

type Board struct {
	Admin       []*Card
	Cards       []core.Card
	IsSpyMaster bool
	exit        chan struct{}
	done        chan struct{}
}

func NewBoard(isSpyMaster bool) *Board {
	b := &Board{
		IsSpyMaster: isSpyMaster,
		Admin:       make([]*Card, 0),
		exit:        make(chan struct{}),
		done:        make(chan struct{}, 1),
	}

	// set up the admin cards already
	inputCard := NewCard("", CardType(1), true)
	inputCard.paragraph.SetRect(0, 30, 35, 35)
	summaryCard := NewCard(summaryText(""), CardType(1), true)
	summaryCard.paragraph.SetRect(40, 30, 80, 35)

	b.Admin = append(b.Admin, inputCard, summaryCard)

	return b
}

func (b *Board) Done() chan struct{} {
	return b.done
}

func (b *Board) Render(renderc chan core.Card, gameCreated chan string, update chan<- string) {

	if err := ui.Init(); err != nil {
		panic(err)
	}

	defer func() {
		ui.Close()
		close(b.done)
	}()

	var keys string

	// render the admin cards before anything else
	for _, c := range b.Admin {
		ui.Render(c.paragraph)
	}

	inputCard := b.Admin[0]
	summaryCard := b.Admin[1]

	for {
		select {
		case id := <-gameCreated:
			summaryCard.paragraph.Text = summaryText(id)
			ui.Render(summaryCard.paragraph)
		case c := <-renderc:
			card := NewCard(c.GetContent(), CardType(c.GetType()), c.IsActive() || b.IsSpyMaster)
			coords := c.Coords()
			card.paragraph.SetRect(int(coords.X1), int(coords.Y1), int(coords.X2), int(coords.Y2))
			ui.Render(card.paragraph)
		case e := <-ui.PollEvents():
			if e.Type == ui.KeyboardEvent {
				if e.ID == "<Enter>" {
					update <- keys
					keys = ""
					inputCard.paragraph.Text = keys
					ui.Render(inputCard.paragraph)
					continue
				}
				if e.ID == "<Backspace>" {
					keys = keys[:len(keys)-1]
					inputCard.paragraph.Text = keys
					ui.Render(inputCard.paragraph)
				} else {
					keys += e.ID
					inputCard.paragraph.Text = keys
					ui.Render(inputCard.paragraph)
				}

				if keys == "!q" {
					return
				}
			}

		}
	}
}

func summaryText(id string) string {
	return fmt.Sprintf("Game ID: %s\nTo Exit Enter: !q\nTo Submit: Type word and hit enter", id)
}
