package storage

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameSavesStorage interface {
	Create(gameSave types.GameSaves) (err error)
	FindByID(gameID uuid.UUID, userID uuid.UUID) (gameSave types.GameSaves, err error)
	Update(gameSave types.GameSaves) (err error)
	Delete(gameID uuid.UUID, userID uuid.UUID) (err error)
}

type gameSavesStorage struct {
	collection *mongo.Collection
}

func (d *gameSavesStorage) Create(gameSave types.GameSaves) (err error) {
	_, err = d.collection.InsertOne(context.TODO(), gameSave)
	return err
}

func (d *gameSavesStorage) FindByID(gameID uuid.UUID, userID uuid.UUID) (gameSave types.GameSaves, err error) {
	filter := bson.M{
		"gameID": gameID,
		"userID": userID,
	}
	err = d.collection.FindOne(context.TODO(), filter).Decode(&gameSave)
	if err != nil {
		return gameSave, err
	}
	return gameSave, nil
}

func (d *gameSavesStorage) Update(gameSave types.GameSaves) (err error) {
	filter := bson.M{
		"gameID": gameSave.GameID,
		"userID": gameSave.UserID,
	}
	update := bson.M{"$set": gameSave}
	_, err = d.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (d *gameSavesStorage) Delete(gameID uuid.UUID, userID uuid.UUID) (err error) {
	filter := bson.M{
		"gameID": gameID,
		"userID": userID,
	}

	_, err = d.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func NewGameSavesStorage(client *mongo.Client) (GameSavesStorage, error) {
	coll := client.Database(dataBaseGame).Collection(gameSavesCollectionKey)
	_, err := coll.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		},
	)
	if err != nil {
		return &gameSavesStorage{}, err
	}

	return &gameSavesStorage{
		collection: coll,
	}, nil
}
