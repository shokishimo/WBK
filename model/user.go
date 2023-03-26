package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	SessionID string     `json:"sessionid"`
	Fav       []Keyboard `json:"fav"`
	BestKeys  []Keyboard `json:"bestkeys"`
	WorstKeys []Keyboard `json:"worstkeys"`
}

// SaveUser stores the user to the specified collection
func (theUser User) SaveUser(collection *mongo.Collection) error {
	_, err := collection.InsertOne(context.TODO(), theUser)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes the user from the specified collection
func (theUser User) DeleteUser(collection *mongo.Collection) error {
	filter := bson.M{"email": theUser.Email}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
