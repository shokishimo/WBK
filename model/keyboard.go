package model

import (
	"context"
	"fmt"
	"github.com/shokishimo/WhatsTheBestKeyboard/database"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

type Keyboard struct {
	Name        string `json:"name"`
	Make        string `json:"make"`
	Price       string `json:"price"`
	Layout      string `json:"layout"`
	Profile     string `json:"profile"`
	Switch      string `json:"switch"`
	Frame       string `json:"frame"`
	Keycap      string `json:"keycap"`
	Ranking     string `json:"ranking"`
	NumVotes    string `json:"numvotes"`
	Score       string `json:"score"`
	Url         string `json:"url"`
	OriginalUrl string `json:"originalurl"`
	Other       string `json:"other"`
}

func FindKeyboardWithNameAndRanking(name string, rank string) (error, Keyboard) {
	db := database.Connect()
	defer db.Disconnect()
	db.GetAccessKeysToKeyboardsCollection()

	// find keyboard
	var theKeyboard Keyboard
	filter := bson.M{"name": name, "ranking": rank}
	err := db.GetCollection().FindOne(context.TODO(), filter).Decode(&theKeyboard)
	// when the keyboard found
	if err != nil {
		return err, Keyboard{}
	}

	return nil, theKeyboard
}

func GetRanks(topHowMany int) []Keyboard {
	db := database.Connect()
	defer db.Disconnect()
	db.GetAccessKeysToKeyboardsCollection()

	// extract keyboard data from database based on their net ranking
	var keyboards []Keyboard
	var keyboard Keyboard
	for i := 1; i <= topHowMany; i++ {
		var t string = strconv.Itoa(i)
		filter := bson.M{"ranking": t}
		err := db.GetCollection().FindOne(context.TODO(), filter).Decode(&keyboard)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		keyboards = append(keyboards, keyboard)
	}
	return keyboards
}
