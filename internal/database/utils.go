package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func getObjectId(id string) (bson.ObjectID, error) {
	idObj, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return idObj, ErrFailedToProcesID
	}
	return idObj, nil
}

func emptyCollection(collectionName string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := GetCollection(collectionName).Drop(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	coll := Client.Database("Blog_DB").Collection(collectionName)
	return coll
}
