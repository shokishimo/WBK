package model

import (
	"context"
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

func SaveUserToTemporaryUsersCollection(theUser User, collection *mongo.Collection) error {
	// begin insert user
	_, err := collection.InsertOne(context.TODO(), theUser)
	if err != nil {
		return err
	}
	return nil
}
