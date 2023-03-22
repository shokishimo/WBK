package db

import (
	"context"
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
		log.Fatal("No .env file found")
	}
	uri = os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to db")
	}
	return client
}

func Disconnect(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal("Failed to disconnect from db")
	}
}

func GetAccessKeysToKeyboardsCollection(client *mongo.Client) *mongo.Collection {
	godotenv.Load()
	database := os.Getenv("DATABASE")
	keyboardCollection := os.Getenv("COLLECTION_Keyboards")
	if database == "" || keyboardCollection == "" {
		log.Fatal("failed to get access keys to database; Error at `GetAccessKeysToKeyboardsCollection()`")
	}
	collection := client.Database(database).Collection(keyboardCollection)
	return collection
}

func GetAccessKeysToUsersCollection(client *mongo.Client) *mongo.Collection {
	godotenv.Load()
	database := os.Getenv("DATABASE")
	userCollection := os.Getenv("COLLECTION_users")
	if database == "" || userCollection == "" {
		log.Fatal("failed to get access keys to database; Error at `GetAccessKeysToKeyboardsCollection()`")
	}
	collection := client.Database(database).Collection(userCollection)
	return collection
}

func GetAccessKeysToTemporaryUsersCollection(client *mongo.Client) *mongo.Collection {
	godotenv.Load()
	database := os.Getenv("DATABASE")
	temporaryUserCollection := os.Getenv("COLLECTION_temporary_users")
	if database == "" || temporaryUserCollection == "" {
		log.Fatal("failed to get access keys to database; Error at `GetAccessKeysToTemporaryUsersCollection()`")
	}
	collection := client.Database(database).Collection(temporaryUserCollection)
	return collection
}
