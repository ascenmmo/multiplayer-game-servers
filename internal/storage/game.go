package storage

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GameStorage interface {
	CreateGame(game types.Game) (err error)
	FindByID(gameID uuid.UUID) (game types.Game, err error)
	FindByName(name string) (game types.Game, err error)
	FindByOwnerID(ownerID uuid.UUID) (game []types.Game, err error)
	Update(game types.Game) (err error)
}

type gameStorage struct {
	collection *mongo.Collection
}

func NewGameStorage(client *mongo.Client) (GameStorage, error) {
	coll := client.Database(dataBaseGame).Collection(gameCollectionKey)
	_, err := coll.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		},
	)
	if err != nil {
		return &gameStorage{}, err
	}

	return &gameStorage{
		collection: coll,
	}, nil

}

func (d *gameStorage) CreateGame(game types.Game) (err error) {
	_, err = d.collection.InsertOne(context.TODO(), game)
	return err
}

func (d *gameStorage) FindByID(gameID uuid.UUID) (game types.Game, err error) {
	filter := bson.M{
		"_id": gameID,
	}
	err = d.collection.FindOne(context.TODO(), filter).Decode(&game)
	if err != nil {
		return game, err
	}
	return game, nil
}

func (d *gameStorage) FindByOwnerID(ownerID uuid.UUID) (game []types.Game, err error) {
	filter := bson.M{
		"owners": bson.M{
			"$elemMatch": bson.M{
				"$eq": ownerID,
			},
		},
	}
	cur, err := d.collection.Find(context.TODO(), filter)
	if err != nil {
		return game, err
	}
	err = cur.All(context.TODO(), &game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (d *gameStorage) FindByName(name string) (game types.Game, err error) {
	filter := bson.M{
		"name": name,
	}

	err = d.collection.FindOne(context.TODO(), filter).Decode(&game)
	if err != nil {
		return game, err
	}
	return game, nil
}

func (d *gameStorage) Update(game types.Game) (err error) {
	filter := bson.M{
		"_id": game.ID,
	}
	update := bson.M{"$set": game}
	_, err = d.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
