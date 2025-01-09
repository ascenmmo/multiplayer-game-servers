package storage

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/internal/errors"
	defalultdata "github.com/ascenmmo/multiplayer-game-servers/internal/storage/defalult_data"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServersStorage interface {
	CreateServer(servers types.Server) (err error)
	FindByID(serversID uuid.UUID) (servers types.Server, err error)
	FindByIDs(serversIDs []uuid.UUID) (servers []types.Server, err error)
	FindAllByServerType(ownerID uuid.UUID, serverType string) (servers []types.Server, err error)
	FindAll(ownerID uuid.UUID) (servers []types.Server, err error)
	Update(servers types.Server) (err error)
	Delete(serversID uuid.UUID) (err error)
	FindAllServerIDs() (serversIDs []uuid.UUID, err error)
}

type serversStorage struct {
	collection *mongo.Collection
}

func (d *serversStorage) CreateServer(servers types.Server) (err error) {
	_, err = d.collection.InsertOne(context.TODO(), servers)
	return err
}

func (d *serversStorage) FindByID(serversID uuid.UUID) (servers types.Server, err error) {
	filter := bson.M{
		"_id": serversID,
	}

	err = d.collection.FindOne(context.TODO(), filter).Decode(&servers)

	if err != nil {
		return servers, err
	}

	return servers, nil
}
func (d *serversStorage) FindByIDs(serversIDs []uuid.UUID) (servers []types.Server, err error) {
	filter := bson.M{"_id": bson.M{"$in": serversIDs}}
	cur, err := d.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	err = cur.All(context.TODO(), &servers)
	if err != nil {
		return nil, err
	}

	if len(servers) == 0 {
		return nil, errors.ErrServerCreatingRoomAllServesOffError
	}

	return servers, nil
}

func (d *serversStorage) FindAllByServerType(ownerID uuid.UUID, serverType string) (servers []types.Server, err error) {
	filter := bson.M{
		"owners": bson.M{
			"$elemMatch": bson.M{
				"$eq": ownerID,
			},
		},
		"server_type": serverType,
	}

	cur, err := d.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	err = cur.All(context.TODO(), &servers)
	if err != nil {
		return nil, err
	}

	return servers, nil
}

func (d *serversStorage) FindAll(ownerID uuid.UUID) (servers []types.Server, err error) {
	filter := bson.M{
		"owners": bson.M{
			"$elemMatch": bson.M{
				"$eq": ownerID,
			},
		},
	}

	cur, err := d.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	err = cur.All(context.TODO(), &servers)
	if err != nil {
		return nil, err
	}

	return servers, nil
}

func (d *serversStorage) Update(servers types.Server) (err error) {
	filter := bson.M{
		"_id": servers.ID,
	}
	update := bson.M{"$set": servers}
	_, err = d.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (d *serversStorage) FindAllServerIDs() (serversIDs []uuid.UUID, err error) {
	filter := bson.M{}

	findOptions := options.Find().SetProjection(bson.M{"_id": 1})
	cursor, err := d.collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var result struct {
			ID uuid.UUID `bson:"_id"`
		}

		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}

		serversIDs = append(serversIDs, result.ID)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return serversIDs, nil
}

func (d *serversStorage) Delete(serversID uuid.UUID) (err error) {
	filter := bson.M{
		"_id": serversID,
	}

	_, err = d.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (d *serversStorage) addDefaultServers() (err error) {
	servers, err := d.FindAll(DefaultUserID)
	if err != nil {
		return err
	}

	if len(servers) == 0 {
		for _, v := range servers {
			err = d.CreateServer(v)
			if err != nil {
				return err
			}
		}
	}

	for _, v := range defalultdata.AddServers(DefaultUserID) {
		err = d.Update(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewServersStorage(client *mongo.Client) (ServersStorage, error) {
	coll := client.Database(dataBaseServers).Collection(serversCollectionKey)
	_, err := coll.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{{
			Keys:    bson.D{{Key: "address", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
			{
				Keys: bson.D{{Key: "creator_id", Value: 1}},
			},
		},
	)
	if err != nil {
		return &serversStorage{}, err
	}

	newServersStorage := &serversStorage{
		collection: coll,
	}

	err = newServersStorage.addDefaultServers()
	if err != nil {
		return &serversStorage{}, err
	}

	return newServersStorage, nil

}
