package mongoconn

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
)

type Connection struct {
	Connectionstring string   `yaml:"connection"`
	Database         []string `yaml:"database"`
}

func Readyaml(filepath string) Connection {
	var conn Connection

	// Open YAML file
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	// Decode YAML file to struct
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&conn); err != nil {
			log.Println(err.Error())
		}
	}
	fmt.Println(conn.Connectionstring)
	fmt.Println(conn.Database[0])
	fmt.Println(conn.Database[1])
	return conn
}

func ConnectDB() *mongo.Client {
	conn := Readyaml("rahul.yaml")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn.Connectionstring))
	if err != nil {
		log.Fatal(err)
	}
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
func GetCollection(client *mongo.Client, databaseName, collectionName string) *mongo.Collection {
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}
