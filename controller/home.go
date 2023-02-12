package controller

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shokishimo/WhatsTheBestKeyboard/db"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"html/template"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ServePublicHome(w)
}

// ServePublicHome shows the public template home to the browser
func ServePublicHome(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("static/public/home.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	tmpl.Execute(w, nil)
}

func AddSampleData() {
	godotenv.Load()
	database := os.Getenv("DATABASE")
	keyboardCollection := os.Getenv("COLLECTION_Keyboard")
	if database == "" || keyboardCollection == "" {
		fmt.Println("failed to get dtata")
	}
	client := db.Connect()
	// Disconnect from db
	defer db.Disconnect(client)

	data1 := model.Keyboard{
		Name:     "Air 75",
		Make:     "NuPhy",
		Price:    "$120",
		Layout:   "75%",
		Profile:  "low",
		Switch:   "mechanical",
		Material: model.Material{},
		Ranking:  "1",
		Score:    "8.5",
	}
	fmt.Println(data1)
	// begin insert data
	collection := client.Database(database).Collection(keyboardCollection)
	a, err := collection.InsertOne(context.TODO(), data1)
	fmt.Println(a)
	if err != nil {
		fmt.Println("aaa")
		fmt.Println(err.Error())
	}
}
