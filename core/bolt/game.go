package bolt

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/alabianca/codernames/core/bolt/models"
	"github.com/boltdb/bolt"
)

var ErrBucketNotFound = "bucket not found"

type GameDAL struct {
	DB *bolt.DB
	BucketName string
}

// Create accepts a protobuf message, marshals it
// and attempts to insert it in the main bucket.
func (dal *GameDAL) Create(doc interface{}) error {
	id, err := randomHex(8)
	if err != nil {
		return err
	}

	x, ok := doc.(*models.Game)
	if !ok {
		return  errors.New("could not convert to protobuf object")
	}

	x.ID = id

	bts, err := json.Marshal(x)
	if err != nil {
		return err
	}

	err = dal.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dal.BucketName))
		if bucket == nil {
			// bucket not found
			return errors.New(ErrBucketNotFound)
		}

		return bucket.Put([]byte(id), bts)
	})

	if err != nil {
		return err
	}

	return nil
}

func (dal *GameDAL) Get(doc interface{}) error  {
	game, ok := doc.(*models.Game)
	if !ok {
		return errors.New("could not convert to game object")
	}

	var gameBts []byte
	err := dal.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dal.BucketName))
		if bucket == nil {
			return errors.New(ErrBucketNotFound)
		}

		gameBts = bucket.Get([]byte(game.ID))
		return nil
	})

	if err != nil {
		return err
	}

	var out models.Game
	if err := json.Unmarshal(gameBts, &out); err != nil {
		return err
	}

	game.ID = out.ID
	game.Cards = out.Cards
	return nil
}

func (dal *GameDAL) Activate(id string, content string, doc interface{}) error {

	return dal.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dal.BucketName))
		if bucket == nil {
			return errors.New(ErrBucketNotFound)
		}

		gameBts := bucket.Get([]byte(id))
		if err := json.Unmarshal(gameBts, doc); err != nil {
			return err
		}

		game, ok := doc.(*models.Game)
		if !ok {
			return errors.New("could not convert to game object")
		}

		var index = -1
		for i, card := range game.Cards {
			if card.Content == content {
				index = i
				break
			}
		}

		if index > -1 {
			game.Cards[index].Active = true
		}

		gameBts, err := json.Marshal(game)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(id), gameBts)
	})
}

func randomHex(n int) (string, error) {
	bts := make([]byte, n)
	if _, err := rand.Read(bts); err != nil {
		return "", err
	}

	return hex.EncodeToString(bts), nil
}