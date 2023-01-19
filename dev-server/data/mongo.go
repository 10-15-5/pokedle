package data

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var dbName = ""

func StartMongoDB() *mongo.Client {

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	database := os.Getenv("DATABASE")

	if database == "" {
		log.Fatal("You must set your 'DATABASE' environmental variable.")
	} else {
		dbName = database
	}

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")

	return mongoClient
}

func CloseMongoDB() {
	err := mongoClient.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}
