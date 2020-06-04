package mongo

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/alabianca/codernames/core/mongo/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type GameDAL struct {
	Collection *mongo.Collection
}

func (dal *GameDAL) Create(doc interface{}) error {
	res, err := dal.Collection.InsertOne(nil, doc)
	if err != nil {
		return err
	}

	game, ok := doc.(*models.Game)
	if !ok {
		return errors.New("could not convert document to game object")
	}

	game.ID, ok = res.InsertedID.(primitive.ObjectID)
	if !ok {
		return errors.New("could not convert to object id")
	}

	return nil
}

func (dal *GameDAL) Get(doc interface{}) error {
	game, ok := doc.(*models.Game)
	if !ok {
		return errors.New("could not convert to game object")
	}

	if game.ID == primitive.NilObjectID {
		return errors.New("object id is required")
	}

	return dal.Collection.FindOne(nil, bson.D{{"_id", game.ID}}).Decode(doc)
}

func (dal *GameDAL) Activate(id string, content string, doc interface{}) error {
	pid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return dal.Collection.FindOneAndUpdate(
		nil,
		bson.D{{"_id", pid}},
		bson.D{{"$set", bson.D{{"cards.$[elem].active", true}}}},
		options.FindOneAndUpdate().SetArrayFilters(
			options.ArrayFilters{
				Filters: []interface{}{bson.D{{"elem.content", content}}},
			},
		).SetReturnDocument(1),
	).Decode(doc)


}
