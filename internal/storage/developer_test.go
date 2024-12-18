package storage

import (
	"context"
	"fmt"
	"github.com/ascenmmo/multiplayer-game-servers/env"
	"github.com/ascenmmo/multiplayer-game-servers/pkg/multiplayer/types"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestGetDevelopers(t *testing.T) {
	generator, err := tokengenerator.NewTokenGenerator(env.TokenKey)
	if err != nil {
		panic(err)
	}

	collection := NewMongoConnection(env.MongoURL).Database(dataBaseDevelopers).Collection(developersCollectionKey)
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	var dev []types.Developer
	err = cur.All(context.TODO(), &dev)
	if err != nil {
		panic(err)
	}

	for _, v := range dev {
		fmt.Println("____dev___")
		secret, err := generator.ParseSecretHash(v.Password)
		if err != nil {
			panic(err)
		}
		fmt.Println(v.Email, secret)
	}
}
