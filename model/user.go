package model

import (
	"context"
	"github.com/shokishimo/WhatsTheBestKeyboard/database"
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
func (theUser User) SaveUser(db database.DB) error {
	_, err := db.GetCollection().InsertOne(context.TODO(), theUser)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes the user from the specified collection
func (theUser User) DeleteUser(db database.DB) error {
	filter := bson.M{"email": theUser.Email}

	result, err := db.GetCollection().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
