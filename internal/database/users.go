package db

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/tu-usuario/blog-api/internal/constants"
	"github.com/tu-usuario/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func FindUserCredentialsByEmail(ctx context.Context, email string) (*models.UserCredentials, error) {

	filter := bson.M{"email": email}

	userCreds, err := findOneUserWithFilter(ctx, filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserDoesNotExist
		}
		return nil, err
	}

	return userCreds, nil
}

func InsertUserCredentials(ctx context.Context, creds *models.UserCredentials) (bson.ObjectID, error) {

	id, err := insertOneIntoCollection(ctx, creds, constants.AuthCredentialsCollections)
	if err != nil {
		switch {
		case errors.Is(err, ErrFailedToInsert):
			return bson.NilObjectID, ErrFailedToInsertUser
		case errors.Is(err, ErrDuplicatedKey):
			return bson.NilObjectID, ErrDuplicatedKeyForUser
		default:
			return bson.NilObjectID, err
		}
	}

	return id, nil
}

func FindAllUsers(ctx context.Context, c *gin.Context) (*[]models.UserCredentials, error) {

	usersCreds := []models.UserCredentials{}
	err := findWithFilterFromCollectionOfType(ctx, bson.M{}, constants.AuthCredentialsCollections, &usersCreds)

	if err != nil {
		switch {
		case errors.Is(err, ErrFailedToFind):
			return nil, ErrFailedToFindUser
		case errors.Is(err, ErrFailedToProccess):
			return nil, ErrFailedToProccessUser
		default:
			return nil, err
		}
	}

	return &usersCreds, nil
}

func findOneUserWithFilter(ctx context.Context, filter bson.M) (*models.UserCredentials, error) {
	var userCreds models.UserCredentials
	err := findOneWithFilterFromColletionOfType(ctx, filter, constants.AuthCredentialsCollections, &userCreds)

	if err != nil {
		switch {
		case errors.Is(err, ErrDoesNotExist):
			err = ErrUserDoesNotExist
		default:
			err = ErrFailedToInsertUser
		}
		return nil, err
	}

	return &userCreds, err
}
