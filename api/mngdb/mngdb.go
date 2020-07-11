package mngdb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Clonodb *mongo.Database
var Ctx = context.TODO()

func InitDB() *mongo.Client {
	var err error
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	Client, err = mongo.Connect(Ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	Clonodb = Client.Database("clon-o-gram")
	fmt.Println("Connected to MongoDB!")
	return Client
}
