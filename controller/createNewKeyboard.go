package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shokishimo/WhatsTheBestKeyboard/db"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"net/http"
	"os"
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

	godotenv.Load()
	database := os.Getenv("DATABASE")
	keyboardCollection := os.Getenv("COLLECTION_Keyboard")
	if database == "" || keyboardCollection == "" {
		fmt.Println("failed to get access keys to database")
	}
	client := db.Connect()
	// Disconnect from db
	defer db.Disconnect(client)

	// begin insert data
	collection := client.Database(database).Collection(keyboardCollection)
	a, err := collection.InsertOne(context.TODO(), keyboard)
	fmt.Println(a)
	if err != nil {
		fmt.Println(err.Error())
	}
}
