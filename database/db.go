package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Db struct {
	clientDB *mongo.Client
}

func Connect() (Db, error) {
	//create empty context
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return Db{}, err
	}

	//create ping
	client.Ping(ctx, readpref.Primary())
	if err != nil {
		return Db{}, nil
	}

	return Db{
		clientDB: client,
	}, nil
}

func (db *Db) GetUserCollection() *mongo.Collection{
	userCollection := db.clientDB.Database("golang").Collection("users")
	return userCollection
}
