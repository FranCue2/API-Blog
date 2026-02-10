package db

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/tu-usuario/blog-api/internal/constants"
	"github.com/tu-usuario/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client

func InitDB() {
	connectDB()
	createUniqueIndexes()
}

func connectDB() {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := os.Getenv("MONGO_URI")

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	opts.SetConnectTimeout(10 * time.Second)

	tmp, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	Client = tmp

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	fmt.Println("✅ Conectado a MongoDB exitosamente")
}

func createUniqueIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := GetCollection("auth_credentials")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := coll.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		fmt.Printf("❌ Error creando indices unicos para los emails: %s", err)
	}

	fmt.Println("✅ Indice unico para los emails")

}

func EmptyCollection(collectionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := GetCollection(collectionName).Drop(ctx)
	if err != nil {
		return err
	}

	fmt.Println("🗑️  Base de datos vaciada exitosamente")
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
			return nil, errors.New("usuario no existe")
		}
		return nil, err
	}

	return &userCred, nil
}
