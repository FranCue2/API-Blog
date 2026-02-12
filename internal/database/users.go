package db

import (
	"context"

	"github.com/tu-usuario/blog-api/internal/constants"
	"github.com/tu-usuario/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func FindUserCredentialsByEmail(ctx context.Context, email string) (*models.UserCredentials, error) {

	filter := bson.M{"email": email}

	var userCred models.UserCredentials

	err := GetCollection(constants.AuthCredentialsCollections).FindOne(ctx, filter).Decode(&userCred)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserDoesNotExist
		}
		return nil, err
	}

	return &userCred, nil
}

func InsertCredentials(ctx context.Context, creds *models.UserCredentials) (*mongo.InsertOneResult, error) {
	res, err := GetCollection(constants.AuthCredentialsCollections).InsertOne(ctx, creds)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			err = ErrUserAlreadyExists
		}
		return nil, err
	}

	return res, nil
}
