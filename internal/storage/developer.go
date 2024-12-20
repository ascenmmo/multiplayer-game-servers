package storage

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeveloperStorage interface {
	CreateDeveloper(developer types.Developer) (err error)
	FindByID(developerID uuid.UUID) (developer types.Developer, err error)
	FindByIDS(developersID []uuid.UUID) (developers []types.Developer, err error)
	FindByPassword(email, nickname, password string) (developer types.Developer, err error)
	Update(developer types.Developer) (err error)
}

type developerStorage struct {
	collection *mongo.Collection
}

func NewDeveloperStorage(client *mongo.Client) (DeveloperStorage, error) {
	coll := client.Database(dataBaseDevelopers).Collection(developersCollectionKey)
	_, err := coll.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{
			{
				Keys: bson.D{
					{Key: "email", Value: 1},
					{Key: "nickname", Value: 1},
					{Key: "password", Value: 1},
				},
			},
		},
	)
	if err != nil {
		return &developerStorage{}, err
	}

	return &developerStorage{
		collection: coll,
	}, nil

}

func (d *developerStorage) CreateDeveloper(developer types.Developer) (err error) {
	_, err = d.collection.InsertOne(context.TODO(), developer)
	return err
}

func (d *developerStorage) FindByIDS(developersID []uuid.UUID) (developers []types.Developer, err error) {
	filter := bson.M{"_id": bson.M{"$in": developersID}}
	cur, err := d.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	err = cur.All(context.TODO(), &developers)
	if err != nil {
		return nil, err
	}

	return developers, nil
}

func (d *developerStorage) FindByID(developerID uuid.UUID) (developer types.Developer, err error) {
	filter := bson.M{
		"_id": developerID,
	}
	err = d.collection.FindOne(context.TODO(), filter).Decode(&developer)
	if err != nil {
		return developer, err
	}
	return developer, nil
}

func (d *developerStorage) FindByPassword(email, nickname, password string) (developer types.Developer, err error) {
	filter := bson.M{
		"password": password,
	}
	if email != "" {
		filter["email"] = email
	}
	if nickname != "" {
		filter["nickname"] = nickname
	}

	err = d.collection.FindOne(context.TODO(), filter).Decode(&developer)
	if err != nil {
		return developer, err
	}
	return developer, nil
}

func (d *developerStorage) Update(developer types.Developer) (err error) {
	filter := bson.M{
		"_id": developer.ID,
	}
	update := bson.M{"$set": developer}
	_, err = d.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
