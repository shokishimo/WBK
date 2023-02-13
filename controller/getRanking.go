package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shokishimo/WhatsTheBestKeyboard/db"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetRankingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get query parameters
	queryParams := r.URL.Query()
	numberOfData, err := strconv.Atoi(queryParams.Get("number"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// database connection
	godotenv.Load()
	database := os.Getenv("DATABASE")
	keyboardCollection := os.Getenv("COLLECTION_Keyboard")
	if database == "" || keyboardCollection == "" {
		fmt.Println("failed to get access keys to database")
	}
	client := db.Connect()
	// Disconnect from db
	defer db.Disconnect(client)
	collection := client.Database(database).Collection(keyboardCollection)

	// extract keyboard data from database based on their net ranking
	var keyboards []model.Keyboard
	var keyboard model.Keyboard
	for i := 1; i <= numberOfData; i++ {
		var t string = strconv.Itoa(i)
		filter := bson.M{"ranking": t}
		err := collection.FindOne(context.TODO(), filter).Decode(&keyboard)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		keyboards = append(keyboards, keyboard)
	}

	// Return the slice of keyboards as a JSON response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(keyboards)
	if err != nil {
		log.Fatalf("Error encoding response: %v", err)
	}
}
