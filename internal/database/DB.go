package db

import (
	"context"
	"fmt"
	"time"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)
var Client *mongo.Client

func ConnectDB() {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := os.Getenv("MONGO_URI")

	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	opts.SetConnectTimeout(10*time.Second)
	
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
