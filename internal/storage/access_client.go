package storage

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccessClientStorage interface {
	CreateAccessClient(accessClient types.AccessClient) (err error)
	FindByID(accessClientID uuid.UUID) (accessClient types.AccessClient, err error)
	FindByGameID(gameID uuid.UUID) (accessClient types.AccessClient, err error)
	Update(accessClientID types.AccessClient) (err error)
}

type accessClientIDStorage struct {
	collection *mongo.Collection
}

func NewAccessClientStorage(client *mongo.Client) (AccessClientStorage, error) {
	coll := client.Database(dataBaseAccess).Collection(accessClientCollectionKey)
	_, err := coll.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		},
	)
	if err != nil {
		return &accessClientIDStorage{}, err
	}

	return &accessClientIDStorage{
		collection: coll,
	}, nil

}

func (d *accessClientIDStorage) CreateAccessClient(accessClient types.AccessClient) (err error) {
	_, err = d.collection.InsertOne(context.TODO(), accessClient)
	return err
}

func (d *accessClientIDStorage) FindByID(accessClientID uuid.UUID) (accessClient types.AccessClient, err error) {
	filter := bson.M{
		"_id": accessClientID,
	}
	err = d.collection.FindOne(context.TODO(), filter).Decode(&accessClient)
	if err != nil {
		return accessClient, err
	}
	return accessClient, nil
}

func (d *accessClientIDStorage) FindByGameID(gameID uuid.UUID) (accessClient types.AccessClient, err error) {
	filter := bson.M{
		"developers": bson.M{
			"$elemMatch": gameID,
		},
	}
	cur, err := d.collection.Find(context.TODO(), filter)
	if err != nil {
		return accessClient, err
	}
	err = cur.All(context.TODO(), &accessClient)
	if err != nil {
		return accessClient, err
	}

	return accessClient, nil
}

func (d *accessClientIDStorage) Update(accessClient types.AccessClient) (err error) {
	filter := bson.M{
		"_id": accessClient.ID,
	}
	update := bson.M{"$set": accessClient}
	_, err = d.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
