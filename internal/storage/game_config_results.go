package storage

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameConfigsResultsStorage interface {
	Create(result types.GameConfigResults) (err error)
	CreateMany(result []types.GameConfigResults) (err error)
	GetAllResults(gameID uuid.UUID, offset, limit int64) (gameConfig []types.GameConfigResults, err error)
	GetResults(gameID uuid.UUID, roomID uuid.UUID) (gameConfig []types.GameConfigResults, err error)
}

type gameConfigsResultsStorage struct {
	collection *mongo.Collection
}

func (s *gameConfigsResultsStorage) Create(result types.GameConfigResults) (err error) {
	_, err = s.collection.InsertOne(context.Background(), result)
	if err != nil {
		return err
	}
	return err
}

func (s *gameConfigsResultsStorage) CreateMany(result []types.GameConfigResults) (err error) {
	docs := make([]interface{}, len(result))
	for i, v := range result {
		docs[i] = v
	}

	_, err = s.collection.InsertMany(context.Background(), docs)
	if err != nil {
		return err
	}
	return err
}

func (s *gameConfigsResultsStorage) GetAllResults(gameID uuid.UUID, offset, limit int64) (gameConfig []types.GameConfigResults, err error) {
	filter := bson.M{
		"gameID": gameID,
	}

	opts := options.Find().SetSort(limit).SetSkip(offset)

	cur, err := s.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return gameConfig, err
	}

	err = cur.All(context.Background(), gameConfig)
	if err != nil {
		return gameConfig, err
	}
	return gameConfig, nil
}

func (s *gameConfigsResultsStorage) GetResults(gameID uuid.UUID, roomID uuid.UUID) (gameConfig []types.GameConfigResults, err error) {
	filter := bson.M{
		"gameID": gameID,
		"roomID": roomID,
	}
	err = s.collection.FindOne(context.Background(), filter).Decode(gameConfig)
	if err != nil {
		return gameConfig, err
	}

	return gameConfig, nil
}

func NewGameConfigsResultsStorage(client *mongo.Client) (GameConfigsResultsStorage, error) {
	coll := client.Database(dataBaseGame).Collection(roomsCollectionKey)
	_, err := coll.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{
			{
				Keys: bson.D{{Key: "gameID", Value: 1}},
			},
			{
				Keys: bson.D{{Key: "roomID", Value: 1}},
			},
		},
	)
	if err != nil {
		return nil, err
	}
	gameConf := gameConfigsResultsStorage{
		collection: coll,
	}
	return &gameConf, nil
}
