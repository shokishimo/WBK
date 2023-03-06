package controller

import (
	"context"
	"github.com/shokishimo/WhatsTheBestKeyboard/db"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"net/http"
)

type LoginData struct {
	ErrorString string
}

func renderLoginPage(w http.ResponseWriter, data LoginData) {
	tmpl, err := template.ParseFiles("static/public/login.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error: Failed to get login page"))
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error: Failed to render login page"))
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderLoginPage(w, LoginData{})
	} else if r.Method == http.MethodPost {
		res := handleLogin(w, r)
		if res != "" {
			renderLoginPage(w, LoginData{ErrorString: res})
		}
		// Redirect to account home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) string {
	email := r.FormValue("email")
	password := r.FormValue("password")
	result := ValidateSignupInput(email, password)
	if result != "" {
		w.WriteHeader(http.StatusBadRequest)
		return result
	}

	// check in database
	client := db.Connect()
	defer db.Disconnect(client)
	collection := db.GetAccessKeysToUsersCollection(client)

	// create a new sessionID
	sessionID := GenerateSessionID()

	var res model.User
	// Define opt to return the updated document
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.M{"email": email, "password": Hash(password)}
	update := bson.M{"$set": bson.M{"sessionid": Hash(sessionID)}}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(&res)
	if err != nil {
		return "Error happened during some executions to database: " + err.Error()
	}

	// when found
	SetUsernameCookie(w, res.Username)
	SetSessionCookie(w, sessionID)

	return ""
}
