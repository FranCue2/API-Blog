package db

import (
	"context"

	"os"
	"time"

	"github.com/tu-usuario/blog-api/internal/constants"
	"github.com/tu-usuario/blog-api/internal/models"
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
		return ErrFailledToCreateIndexis
	}

	return nil

}

func EmptyCollection(collectionName string) error {

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
