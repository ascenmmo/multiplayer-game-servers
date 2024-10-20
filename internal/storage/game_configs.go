package storage

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameConfigsStorage interface {
	CreateOrUpdateNewConfig(gameConfig types.GameConfigs) (err error)
	GetConfig(gameID uuid.UUID) (gameConfig types.GameConfigs, err error)
}

type gameConfigsStorage struct {
	collection *mongo.Collection
}

func (s *gameConfigsStorage) CreateOrUpdateNewConfig(gameConfig types.GameConfigs) (err error) {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"_id": gameConfig.GameID,
	}
	update := bson.M{"$set": gameConfig}
	_, err = s.collection.UpdateOne(context.TODO(), filter, update, opts)
	return err
}

func (s *gameConfigsStorage) GetConfig(gameID uuid.UUID) (gameConfig types.GameConfigs, err error) {
	filter := bson.M{
		"_id": gameID,
	}

	err = s.collection.FindOne(context.Background(), filter).Decode(&gameConfig)
	if err != nil {
		return gameConfig, err
	}
	return gameConfig, nil
}

func NewGameConfigsStorage(client *mongo.Client) (GameConfigsStorage, error) {
	coll := client.Database(dataBaseGame).Collection(configsCollectionKey)

	gameConf := gameConfigsStorage{
		collection: coll,
	}
	return &gameConf, nil
}
