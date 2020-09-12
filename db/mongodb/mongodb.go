package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoClient *mongo.Client
)

func init() {
	var err error

	MongoClient, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("fail trying to create new client for mongodb")
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = MongoClient.Connect(ctx)
	if err != nil {
		fmt.Println("fail trying to connect to mongodb")
		panic(err)
	}
	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("fail trying to Ping to mongodb")
		panic(err)
	}
	fmt.Println("mongodb connected!")
}
