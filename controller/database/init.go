package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func ConnDB() {
	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	db = client.Database("users")
	fmt.Println("Successfuly connected to the database.")
}

// func DBInit() {
// 	fmt.Println("Starting mongoDB...")
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		print("ERROR: ", err)
// 	}
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Connected to MongoDB!")
// }
