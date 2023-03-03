package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shokishimo/WhatsTheBestKeyboard/db"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"net/http"
)

func CreateNewKeyboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var err error
	decoder := json.NewDecoder(r.Body)
	var keyboard model.Keyboard
	err = decoder.Decode(&keyboard)
	if err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusBadRequest)
		return
	}

	client := db.Connect()
	defer db.Disconnect(client)
	// Obtain collection
	collection := db.GetAccessKeysToKeyboardsCollection(client)

	// begin insert data
	_, err = collection.InsertOne(context.TODO(), keyboard)

	if err != nil {
		fmt.Println(err.Error())
	}
}
