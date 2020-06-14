package server

import (
	"context"
	"fmt"
	"github.com/alabianca/codernames/core"
	"github.com/alabianca/codernames/core/bolt/models"
	"github.com/alabianca/codernames/core/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var protoCardTypeMapping = map[pb.Card_Type]models.CardType{
	pb.Card_RED:      models.RED,
	pb.Card_BLUE:     models.BLUE,
	pb.Card_ASSASSIN: models.ASSASSIN,
	pb.Card_CIVILIAN: models.CIVILIAN,
}

var modelCardTypeMapping = map[models.CardType]pb.Card_Type{
	models.RED:      pb.Card_RED,
	models.BLUE:     pb.Card_BLUE,
	models.CIVILIAN: pb.Card_CIVILIAN,
	models.ASSASSIN: pb.Card_ASSASSIN,
}

type Server struct {
	Game   core.GameDAL
	pubsub *PubSub
}

func NewGameService(dal core.GameDAL) *Server {
	return &Server{Game: dal, pubsub: NewPubSub()}
}

func (s *Server) CreateGame(ctx context.Context, request *pb.CreateGameRequest) (*pb.CreateGameResponse, error) {
	var game models.Game

	for _, c := range request.Game.GetCards() {
		game.Cards = append(game.Cards, models.Card{
			X1:       c.GetX1(),
			X2:       c.GetX2(),
			Y1:       c.GetY1(),
			Y2:       c.GetY2(),
			Content:  c.GetContent(),
			CardType: protoCardTypeMapping[c.Type],
			Active:   false,
		})
	}

	if err := s.Game.Create(&game); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error creating the game %v", err))
	}

	var res pb.CreateGameResponse
	var createdGame pb.Game

	createdGame.Id = game.ID
	for _, card := range game.Cards {
		c := &pb.Card{
			X1:      card.X1,
			X2:      card.X2,
			Y1:      card.Y1,
			Y2:      card.Y2,
			Content: card.Content,
			Type:    modelCardTypeMapping[card.CardType],
		}

		createdGame.Cards = append(createdGame.Cards, c)
	}

	res.Game = &createdGame

	return &res, nil
}

func (s *Server) JoinGame(req *pb.JoinGameRequest, stream pb.GameService_JoinGameServer) error {
	var game models.Game

	game.ID = req.GetId()

	err := s.Game.Get(&game)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Could not find game %v", err))
	}

	// send the initial message over the stream to render the first state of the game
	s.updateStream(stream, game)
	// subscribe to the game to receive updates
	for g := range s.pubsub.Subscribe(req.Id) {
		fmt.Println("Sending an update stream message")
		s.updateStream(stream, g)
	}

	return nil

}

func (s *Server) UpdateGame(ctx context.Context, req *pb.UpdateGameRequest) (*pb.UpdateGameResponse, error) {
	var game models.Game
	fmt.Println("Updating ", req.GetContent())
	if err := s.Game.Activate(req.GetId(), req.GetContent(), &game); err != nil {
		return nil, err
	}

	var res pb.UpdateGameResponse
	res.Id = game.ID
	// publish the new game state to all subscribed clients
	s.pubsub.Publish(res.Id, game)

	return &res, nil
}

func (s *Server) updateStream(stream pb.GameService_JoinGameServer, game models.Game) {
	if err := stream.Send(makeGameResponse(game)); err != nil {
		fmt.Println("Error sending an update message into the stream")
	}

}

func makeGameResponse(inGame models.Game) *pb.JoinGameResponse {
	var game pb.Game
	game.Id = inGame.ID
	for _, card := range inGame.Cards {
		c := pb.Card{
			Content: card.Content,
			X1:      card.X1,
			X2:      card.X2,
			Y1:      card.Y1,
			Y2:      card.Y2,
			Type:    modelCardTypeMapping[card.CardType],
			Active:  card.Active,
		}

		game.Cards = append(game.Cards, &c)
	}

	return &pb.JoinGameResponse{
		Game: &game,
	}
}
