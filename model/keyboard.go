package model

import (
	"context"
	"fmt"
	"github.com/shokishimo/WhatsTheBestKeyboard/db"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

type Keyboard struct {
	Name     string `json:"name"`
	Make     string `json:"make"`
	Price    string `json:"price"`
	Layout   string `json:"layout"`
	Profile  string `json:"profile"`
	Switch   string `json:"switch"`
	Frame    string `json:"frame"`
	Keycap   string `json:"keycap"`
	Ranking  string `json:"ranking"`
	NumVotes string `json:"numvotes"`
	Score    string `json:"score"`
	Url      string `json:"url"`
	Other    string `json:"other"`
}

func GetRanks(topHowMany int) []Keyboard {
	client := db.Connect()
	defer db.Disconnect(client)
	// Obtain collection
	collection := db.GetAccessKeysToKeyboardsCollection(client)

	// extract keyboard data from database based on their net ranking
	var keyboards []Keyboard
	var keyboard Keyboard
	for i := 1; i <= topHowMany; i++ {
		var t string = strconv.Itoa(i)
		filter := bson.M{"ranking": t}
		err := collection.FindOne(context.TODO(), filter).Decode(&keyboard)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		keyboards = append(keyboards, keyboard)
	}
	return keyboards
}
