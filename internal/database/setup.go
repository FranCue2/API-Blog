package db

import (
	"context"

	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client

func InitDB() error {
	err := connectDB()

	if err != nil {
		return err
	}

	err = createUniqueIndexes()

	return err
}

func connectDB() error {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := os.Getenv("MONGO_URI")

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	opts.SetConnectTimeout(10 * time.Second)

	tmp, err := mongo.Connect(opts)
	if err != nil {
		return ErrFailedToConnectToDataBase
	}

	Client = tmp

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Client.Ping(ctx, nil); err != nil {
		return ErrFailedToConnectToDataBase
	}

	return nil
}

func createUniqueIndexes() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := GetCollection("auth_credentials")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := coll.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return ErrFailedToCreateIndexis
	}

	return nil

}