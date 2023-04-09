package controller

import (
	"context"
	"encoding/json"
	"github.com/shokishimo/WhatsTheBestKeyboard/database"
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
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Error parsing JSON data: " + err.Error()))
		return
	}

	db := database.Connect()
	defer db.Disconnect()
	db.GetAccessKeysToKeyboardsCollection()

	// begin insert data
	_, err = db.GetCollection().InsertOne(context.TODO(), keyboard)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}
