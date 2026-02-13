package db

import (
	"context"
	"errors"
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

func findWithFilterFromCollectionOfType[T any](ctx context.Context, filter bson.M, collection string, listOfResults *[]T) error {
	res, err := GetCollection(collection).Find(ctx, filter)

	if err != nil {
		return ErrFailedToFind
	}

	if err := res.All(ctx, listOfResults); err != nil {
		return ErrFailedToProccess
	}

	return nil
}

func findOneWithFilterFromColletionOfType[T any](ctx context.Context, filter bson.M, collection string, result *T) error {

	res := GetCollection(collection).FindOne(ctx, filter)

	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments){
			return ErrDoesNotExist
		}
		return ErrFailedToFind
	}

	if err := res.Decode(result); err != nil {
		return ErrFailedToProccessPosts
	}

	return nil
}

func deleteOneWithFilterFromCollection(ctx context.Context, filter bson.M, collection string) error {
	_, err := GetCollection(collection).DeleteOne(ctx, filter)
	return err
}

func insertOneIntoCollection[T any](ctx context.Context, obj T, collection string) (bson.ObjectID, error) {
	res, err := GetCollection(collection).InsertOne(ctx, obj)

	if err != nil {
		
		switch{
		case mongo.IsDuplicateKeyError(err):
			return bson.NilObjectID, ErrDuplicatedKey
		default:
			return bson.NilObjectID, ErrFailedToInsert
		}
	}

	id, ok := res.InsertedID.(bson.ObjectID)

	if !ok {
		return bson.NilObjectID, ErrFailedToObtainID
	}
	return id, nil
}