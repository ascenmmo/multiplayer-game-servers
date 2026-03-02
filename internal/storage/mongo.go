package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/ascenmmo/multiplayer-game-servers/env"
	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DefaultUserID = uuid.NewSHA1(uuid.UUID{}, []byte(env.ServerAddress))
)

const (
	dataBaseClients      = "clients"
	clientsCollectionKey = "clients"

	dataBaseDevelopers      = "developers"
	developersCollectionKey = "developers"

	dataBaseGame               = "games"
	gameCollectionKey          = "games"
	roomsCollectionKey         = "rooms"
	configsCollectionKey       = "configs"
	configsResultCollectionKey = "configs_results"
	gameSavesCollectionKey     = "games_saves"

	dataBaseAccess            = "access"
	accessGameCollectionKey   = "accessGame"
	accessClientCollectionKey = "accessClient"

	dataBaseServers      = "servers"
	serversCollectionKey = "servers"
)

func NewMongoConnection(cfg string) *mongo.Client {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg))
	if err != nil {
		log.Fatal(fmt.Sprintf("Error connecting to Mongo. You set need correct mongo url: %v your url is %s", err, cfg))
	}

	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(fmt.Sprintf("Error get info frm Mongo. You set need correct mongo url: %v your url is %s", err, cfg))
	}

	return mongoClient
}
