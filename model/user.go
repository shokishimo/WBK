package model

import (
	"context"
	"github.com/shokishimo/WhatsTheBestKeyboard/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	SessionID1 string     `json:"sessionid1"`
	SessionID2 string     `json:"sessionid2"`
	SessionID3 string     `json:"sessionid3"`
	Fav        []Keyboard `json:"fav"`
	BestKeys   []Keyboard `json:"bestkeys"`
	WorstKeys  []Keyboard `json:"worstkeys"`
}

func CreatNewUser(username string, email string, password string) User {
	return User{
		Username:   username,
		Email:      email,
		Password:   password,
		SessionID1: "",
		SessionID2: "",
		SessionID3: "",
		Fav:        []Keyboard{},
		BestKeys:   []Keyboard{},
		WorstKeys:  []Keyboard{},
	}
}

// SaveUser stores the user to the specified collection
func (theUser *User) SaveUser() error {
	db := database.Connect()
	defer db.Disconnect()
	db.GetAccessKeysToUsersCollection()
	_, err := db.GetCollection().InsertOne(context.TODO(), *theUser)
	if err != nil {
		return err
	}
	return nil
}

// SaveUserToTemporary stores the user to the temporary collection
func (theUser *User) SaveUserToTemporary() error {
	db := database.Connect()
	defer db.Disconnect()
	db.GetAccessKeysToTemporaryUsersCollection()
	_, err := db.GetCollection().InsertOne(context.TODO(), *theUser)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserFromTemporary deletes the user from the temporary collection
func (theUser *User) DeleteUserFromTemporary() error {
	db := database.Connect()
	defer db.Disconnect()
	db.GetAccessKeysToTemporaryUsersCollection()
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

// FindUserWithEmail searches in the database if the user with the email exists
func FindUserWithEmail(email string) (User, error) {
	db := database.Connect()
	defer db.Disconnect()
	db.GetAccessKeysToUsersCollection()
	// check if the input user already exists in the database
	// Define the filter to find a specific document
	var res User = User{}
	filter := bson.M{"email": email}
	err := db.GetCollection().FindOne(context.TODO(), filter).Decode(&res)
	// when the user with the sessionID found
	if err == nil {
		return res, err
	}
	return res, err
}

// FindUserWithPasscode is used to validate the user's passcode when signing up
func FindUserWithPasscode(inPasscode string) (User, error) {
	db := database.Connect()
	defer db.Disconnect()
	db.GetAccessKeysToTemporaryUsersCollection()
	// check if the input user already exists in the database
	// Define the filter to find a specific document
	var theUser User
	filter := bson.M{"sessionid1": inPasscode}
	err := db.GetCollection().FindOne(context.TODO(), filter).Decode(&theUser)
	// when the user with the passcode not found
	if err != nil {
		return User{}, err
	}
	return theUser, nil
}
