package model

import (
	"context"
	"github.com/shokishimo/WhatsTheBestKeyboard/controller"
	"github.com/shokishimo/WhatsTheBestKeyboard/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	Password        string     `json:"password"`
	SessionID       []string   `json:"sessionid"`
	ActiveUserCount int        `json:"activeusercount"`
	MaxActiveUsers  int        `json:"maxactiveUsers"`
	Fav             []Keyboard `json:"fav"`
	BestKeys        []Keyboard `json:"bestkeys"`
	WorstKeys       []Keyboard `json:"worstkeys"`
}

func CreatNewUser(username string, email string, password string) User {
	return User{
		Username:        username,
		Email:           email,
		Password:        password,
		SessionID:       []string{controller.GeneratePasscode()},
		ActiveUserCount: 1,
		MaxActiveUsers:  3,
		Fav:             []Keyboard{},
		BestKeys:        []Keyboard{},
		WorstKeys:       []Keyboard{},
	}
}

// SaveUser stores the user to the specified collection
func (theUser User) SaveUser() error {
	db := database.Connect()
	defer db.Disconnect()
	db.GetAccessKeysToUsersCollection()
	_, err := db.GetCollection().InsertOne(context.TODO(), theUser)
	if err != nil {
		return err
	}
	return nil
}

// SaveUserToTemporary stores the user to the temporary collection
func (theUser User) SaveUserToTemporary() error {
	db := database.Connect()
	defer db.Disconnect()
	db.GetAccessKeysToTemporaryUsersCollection()
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

func FindUserWithEmail(email string) (User, error) {
	db := database.Connect()
	defer db.Disconnect()
	db = db.GetAccessKeysToUsersCollection()

	// check if the input user already exists in the database
	// Define the filter to find a specific document
	var res User
	filter := bson.M{"email": email}
	err := db.GetCollection().FindOne(context.TODO(), filter).Decode(&res)
	// when the user with the sessionID found
	if err != nil {
		return User{}, err
	}

	return res, nil
}
