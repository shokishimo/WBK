package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type DB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (db DB) GetClient() *mongo.Client {
	return db.client
}

func (db DB) GetCollection() *mongo.Collection {
	return db.collection
}

func Connect() DB {
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
		log.Fatal("Failed to connect to database")
	}
	return DB{client: client}
}

func (db DB) Disconnect() {
	if err := db.client.Disconnect(context.TODO()); err != nil {
		log.Fatal("Failed to disconnect from database")
	}
}

func (db DB) GetAccessKeysToKeyboardsCollection() DB {
	return db.getAccessKeys("COLLECTION_Keyboards")
}

func (db DB) GetAccessKeysToUsersCollection() DB {
	return db.getAccessKeys("COLLECTION_users")
}

func (db DB) GetAccessKeysToTemporaryUsersCollection() DB {
	return db.getAccessKeys("COLLECTION_temporary_users")
}

func (db DB) getAccessKeys(collection string) DB {
	godotenv.Load()
	databaseName := os.Getenv("DATABASE")
	collectionName := os.Getenv(collection)
	if databaseName == "" || collectionName == "" {
		log.Fatal("failed to get access keys to database; Error at `GetAccessKeys()`")
	}
	db.collection = db.client.Database(databaseName).Collection(collectionName)
	return db
}
