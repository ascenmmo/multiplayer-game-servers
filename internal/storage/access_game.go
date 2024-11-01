package storage

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccessGameStorage interface {
	CreateAccessGame(accessGame types.AccessGame) (err error)
	FindByID(accessGameID uuid.UUID) (accessGame types.AccessGame, err error)
	FindByGameID(gameID uuid.UUID) (accessGame types.AccessGame, err error)
	FindByOwnerID(ownerID uuid.UUID) (accessGame []types.AccessGame, err error)
	Update(accessGame types.AccessGame) (err error)
	RemoveByGameID(gameID uuid.UUID, creatorID uuid.UUID) (err error)
}

type accessGameStorage struct {
	collection *mongo.Collection
}

func NewAccessGameStorage(client *mongo.Client) (AccessGameStorage, error) {
	coll := client.Database(dataBaseAccess).Collection(accessGameCollectionKey)
	_, err := coll.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{{
			Keys:    bson.D{{Key: "gameID", Value: 1}},
			Options: options.Index().SetUnique(true),
		}, {
			Keys: bson.D{{Key: "owners", Value: 1}},
		},
		},
	)
	if err != nil {
		return &accessGameStorage{}, err
	}

	return &accessGameStorage{
		collection: coll,
	}, nil

}

func (d *accessGameStorage) CreateAccessGame(accessGame types.AccessGame) (err error) {
	_, err = d.collection.InsertOne(context.TODO(), accessGame)
	return err
}

func (d *accessGameStorage) FindByID(accessGameID uuid.UUID) (accessGame types.AccessGame, err error) {
	filter := bson.M{
		"_id": accessGameID,
	}
	err = d.collection.FindOne(context.TODO(), filter).Decode(&accessGame)
	if err != nil {
		return accessGame, err
	}
	return accessGame, nil
}

func (d *accessGameStorage) FindByGameID(gameID uuid.UUID) (accessGame types.AccessGame, err error) {
	filter := bson.M{"gameID": gameID}
	err = d.collection.FindOne(context.TODO(), filter).Decode(&accessGame)
	if err != nil {
		return accessGame, err
	}

	return accessGame, nil
}

func (d *accessGameStorage) FindByOwnerID(ownersID uuid.UUID) (accessGame []types.AccessGame, err error) {
	filter := bson.M{
		"owners": bson.M{
			"$elemMatch": ownersID,
		},
	}
	cur, err := d.collection.Find(context.TODO(), filter)
	if err != nil {
		return accessGame, err
	}
	err = cur.All(context.TODO(), &accessGame)
	if err != nil {
		return accessGame, err
	}

	return accessGame, nil
}

func (d *accessGameStorage) Update(accessGame types.AccessGame) (err error) {
	filter := bson.M{
		"_id": accessGame.ID,
	}
	update := bson.M{"$set": accessGame}
	_, err = d.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (d *accessGameStorage) RemoveByGameID(gameID uuid.UUID, creatorID uuid.UUID) (err error) {
	filter := bson.M{"gameID": gameID, "creator_id": creatorID}
	_, err = d.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
