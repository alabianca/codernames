package client

import (
	"context"
	"fmt"
	"github.com/alabianca/codernames/core"
	"github.com/alabianca/codernames/core/pb"
	"io"
	"log"
	"os"
)

type Opt = func(c *Client)

var cardTypeMap = map[int]pb.Card_Type{
	red:   pb.Card_RED,
	blue:  pb.Card_BLUE,
	white: pb.Card_CIVILIAN,
	black: pb.Card_ASSASSIN,
}

var pbCardTypeMap = map[pb.Card_Type]int{
	pb.Card_CIVILIAN: white,
	pb.Card_ASSASSIN: black,
	pb.Card_RED:      red,
	pb.Card_BLUE:     blue,
}

// CreateGame is an Option that configures the Client as the user that initially created the game
func CreateGame() Opt {
	return func(c *Client) {
		c.Create = true
		c.IsSpyMaster = true
	}
}

// Source is an option that provides a path to a file
// the file is opened and used as the source of words that can be used
func Source(src string) Opt {
	return func(c *Client) {
		c.SrcReader, _ = os.Open(src)
	}
}
// SourceReader is just like Source except that reader should contain words for this game
func SourceReader(reader io.Reader) Opt {
	return func(c *Client) {
		c.SrcReader = reader
	}
}
// SpyMaster configures this client as a spyMaster
func SpyMaster() Opt {
	return func(c *Client) {
		c.IsSpyMaster = true
	}
}

//JoinId configures this client as a regular user joining a game with id.
// if isSpyMaster is set to true, the client is als configured as the spyMaster
func JoinId(id string, isSpyMaster bool) Opt {
	return func(c *Client) {
		c.JoinId = id
		c.IsSpyMaster = isSpyMaster
	}
}
// Client represents a user in the game
type Client struct {
	GRPCClient  pb.GameServiceClient
	Create      bool
	IsSpyMaster bool
	SrcReader   io.Reader
	JoinId      string
}

func NewClient(sc pb.GameServiceClient, opts ...Opt) *Client {
	var c Client
	c.GRPCClient = sc

	for _, opt := range opts {
		opt(&c)
	}

	return &c
}

// Run initializes the game and loops until `done` is closed.
// if Client.Create is true, we first create the game before joining it.
func (c *Client) Run(ctx context.Context, render chan<- core.Card, keys <-chan string, created chan<- string, done chan struct{}) {
	var id string
	if c.Create {
		var err error
		id, err = c.createGame(ctx)
		if err != nil {
			panic(err)
		}
		c.JoinId = id

	} else {
		id = c.JoinId
	}

	created <- id
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go c.joinGame(cctx, id, render)
	go c.updateGame(cctx, keys)

	<-done
	fmt.Println("FINISHED CLIENT")
}

// createGame creates the game and returns the hex id
func (c *Client) createGame(ctx context.Context) (string, error) {

	words, err := core.ProcessWords(c.SrcReader)
	if err != nil {
		return "", err
	}

	cards, err := generate25Cards(words)
	if err != nil {
		return "", err
	}
	var index int
	var game pb.Game
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			x1 := j * 25
			y1 := i * 5
			x2 := (j + 1) * 25
			y2 := (i + 1) * 5
			var pbCard pb.Card
			pbCard.Content = cards[index].content
			pbCard.X1 = int32(x1)
			pbCard.Y1 = int32(y1)
			pbCard.X2 = int32(x2)
			pbCard.Y2 = int32(y2)
			pbCard.Type = cardTypeMap[cards[index].ctype]
			game.Cards = append(game.Cards, &pbCard)
			index++
		}
	}

	req := pb.CreateGameRequest{
		Game: &game,
	}

	res, err := c.GRPCClient.CreateGame(ctx, &req)
	if err != nil {
		return "", err
	}

	return res.Game.Id, nil
}

// joinGame joins the client to the game and sends card messages to the render channel everytime
// a new message is received from the server. A new message is essentially an updated game state
func (c *Client) joinGame(ctx context.Context, id string, render chan<- core.Card) {
	req := pb.JoinGameRequest{
		Id: id,
	}
	stream, err := c.GRPCClient.JoinGame(ctx, &req)
	if err != nil {
		panic(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Unexpected Error %v", err)
			return
		}

		game := res.GetGame()

		for _, pbCard := range game.GetCards() {
			c := new(card)
			c.ctype = pbCardTypeMap[pbCard.Type]
			c.x1 = pbCard.GetX1()
			c.y1 = pbCard.GetY1()
			c.x2 = pbCard.GetX2()
			c.y2 = pbCard.GetY2()
			c.content = pbCard.GetContent()
			c.isActive = pbCard.GetActive()
			render <- c
		}

	}
}
// updateGame updates the game everytime a word is read from the keyEv channel
func (c *Client) updateGame(ctx context.Context, keyEv <-chan string) {
	for word := range keyEv {
		req := pb.UpdateGameRequest{
			Id:      c.JoinId,
			Content: word,
		}
		if _, err := c.GRPCClient.UpdateGame(ctx, &req); err != nil {
			break
		}
	}
}
