package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectToDB connect to the MongoDB and return the client to make request to the DB
func ConnectToDB(connString string) *mongo.Client {
	var err error

	// "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		fmt.Println("fail trying to create new client for mongodb")
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("fail trying to connect to mongodb")
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("fail trying to Ping to mongodb")
		panic(err)
	}
	fmt.Println("mongodb connected!")
	return client
}
