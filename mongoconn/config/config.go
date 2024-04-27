package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	//client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("bxbsubxs"))
	if err != nil {
		log.Fatal(err)
	}

	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client

}

// client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("user-database").Collection(collectionName)
	return collection
}

func GetDatabase(database string) (string, error) {
	return "New DB", fmt.Errorf("This is not correct")
}
