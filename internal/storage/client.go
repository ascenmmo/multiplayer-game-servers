package storage

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientStorage interface {
	CreateClient(client types.Client) (err error)
	FindByID(clientID uuid.UUID, gameID uuid.UUID) (client types.Client, err error)
	FindByPassword(gameID uuid.UUID, email, nickname, password string) (client types.Client, err error)
	Update(client types.Client) (err error)
}

type clientStorage struct {
	collection *mongo.Collection
}

func (c *clientStorage) CreateClient(client types.Client) (err error) {
	_, err = c.collection.InsertOne(context.TODO(), client)
	return err
}

func (c *clientStorage) FindByID(clientID uuid.UUID, gameID uuid.UUID) (client types.Client, err error) {
	filter := bson.M{
		"_id":    clientID,
		"gameID": gameID,
	}
	err = c.collection.FindOne(context.TODO(), filter).Decode(&client)
	if err != nil {
		return client, err
	}
	return client, nil
}

func (c *clientStorage) FindByPassword(gameID uuid.UUID, email, nickname, password string) (client types.Client, err error) {
	filter := bson.M{
		"gameID":   gameID,
		"password": password,
	}
	if email != "" {
		filter["email"] = email
	}
	if nickname != "" {
		filter["nickname"] = nickname
	}
	err = c.collection.FindOne(context.TODO(), filter).Decode(&client)
	if err != nil {
		return client, err
	}
	return client, nil
}

func (c *clientStorage) Update(client types.Client) (err error) {
	filter := bson.M{
		"_id": client.ID,
	}
	update := bson.M{"$set": client}
	_, err = c.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func NewClientStorage(client *mongo.Client) (ClientStorage, error) {
	coll := client.Database(dataBaseClients).Collection(clientsCollectionKey)
	_, err := coll.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{
			{
				Keys: bson.D{
					{Key: "email", Value: 1},
					{Key: "nickname", Value: 1},
					{Key: "password", Value: 1},
					{Key: "gameID", Value: 1},
				},
			},
		},
	)
	if err != nil {
		return &clientStorage{}, nil
	}

	return &clientStorage{
		collection: coll,
	}, nil

}
