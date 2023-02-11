package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func Connect() *mongo.Client {
	var client *mongo.Client
	var uri string
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return nil
	}
	uri = os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Failed to connect to db")
		return nil
	}
	fmt.Println("Success in connecting to MongoDB")
	return client
}

func Disconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Println("Failed to disconnect from db")
		return
	}
	fmt.Println("Disconnected from MongoDB")
}
