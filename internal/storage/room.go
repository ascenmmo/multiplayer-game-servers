package storage

import (
	"context"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type RoomsStorage interface {
	CreateRoom(room types.Room) (err error)
	FindByID(roomID uuid.UUID) (room types.Room, err error)
	FindByRoomCode(gameId uuid.UUID, roomCode string) (room types.Room, err error)
	FindByCreatorID(creatorID uuid.UUID) (room types.Room, err error)
	FindAll(gameID uuid.UUID) (rooms []types.Room, err error)
	Update(rooms types.Room) (err error)
	Delete(roomID uuid.UUID) (err error)
	GetByTime(searchingTime time.Time, limit int64) (rooms []types.Room, err error)
}

type roomsStorage struct {
	collection *mongo.Collection
}

func NewRoomsStorage(client *mongo.Client) (RoomsStorage, error) {
	coll := client.Database(dataBaseGame).Collection(roomsCollectionKey)
	_, err := coll.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{
			{
				Keys: bson.D{{Key: "creator_id", Value: 1}},
			},
			{
				Keys: bson.D{{Key: "gameID", Value: 1}},
			},
		},
	)
	if err != nil {
		return &roomsStorage{}, err
	}

	return &roomsStorage{
		collection: coll,
	}, nil

}

func (d *roomsStorage) CreateRoom(room types.Room) (err error) {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"_id": room.ID,
	}
	update := bson.M{"$set": room}
	_, err = d.collection.UpdateOne(context.TODO(), filter, update, opts)
	return err
}

func (d *roomsStorage) FindByRoomCode(gameId uuid.UUID, roomCode string) (room types.Room, err error) {
	filter := bson.M{"gameID": gameId, "roomCode": roomCode}
	err = d.collection.FindOne(context.TODO(), filter).Decode(&room)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (d *roomsStorage) FindByID(roomID uuid.UUID) (room types.Room, err error) {
	filter := bson.M{"_id": roomID}
	err = d.collection.FindOne(context.TODO(), filter).Decode(&room)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (d *roomsStorage) FindByCreatorID(creatorID uuid.UUID) (room types.Room, err error) {
	filter := bson.M{"creator_id": creatorID}
	err = d.collection.FindOne(context.TODO(), filter).Decode(&room)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (d *roomsStorage) FindAll(gameID uuid.UUID) (rooms []types.Room, err error) {
	filter := bson.M{
		"gameID": gameID,
	}

	cur, err := d.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	err = cur.All(context.TODO(), &rooms)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (d *roomsStorage) Update(room types.Room) (err error) {
	filter := bson.M{
		"_id": room.ID,
	}
	update := bson.M{"$set": room}
	_, err = d.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (d *roomsStorage) Delete(roomID uuid.UUID) (err error) {
	filter := bson.M{
		"_id": roomID,
	}

	_, err = d.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (d *roomsStorage) GetByTime(searchingTime time.Time, limit int64) (rooms []types.Room, err error) {
	filter := bson.D{{Key: "createdAt", Value: bson.D{{Key: "$lt", Value: searchingTime}}}}
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: 1}}).SetLimit(limit)

	cur, err := d.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return rooms, err
	}

	err = cur.All(context.Background(), &rooms)
	if err != nil {
		return rooms, err
	}

	return rooms, nil
}
