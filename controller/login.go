package controller

import (
	"context"
	"errors"
	"github.com/shokishimo/WhatsTheBestKeyboard/database"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"html/template"
	"net/http"
)

type LoginData struct {
	ErrorString string
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

func handleLogin(w http.ResponseWriter, r *http.Request) string {
	email := r.FormValue("email")
	password := r.FormValue("password")
	result := ValidateSignupInput(email, password)
	if result != "" {
		w.WriteHeader(http.StatusBadRequest)
		return result
	}

	theUser, err, sessionid := LoginSessions(email, password)
	if err != nil {
		return err.Error()
	}

	// when found
	SetUsernameCookie(w, theUser.Username)
	SetSessionCookie(w, sessionid)

	return ""
}

func LoginSessions(email string, password string) (model.User, error, string) {
	db := database.Connect()
	defer db.Disconnect()
	db = db.GetAccessKeysToUsersCollection()

	var theUser model.User
	filter := bson.M{"email": email, "password": Hash(password)}
	err := db.GetCollection().FindOne(context.TODO(), filter).Decode(&theUser)
	// when the user with the passcode not found
	if err != nil {
		return model.User{}, err, ""
	}

	sessionid := GenerateSessionID()
	session := "session"
	// When 3 devices are filled
	if theUser.SessionID1 != "" && theUser.SessionID2 != "" && theUser.SessionID3 != "" {
		return model.User{}, errors.New("one can access their account with up to 3 devices. No more device is available"), ""
	} else if theUser.SessionID1 == "" {
		session += "1"
	} else if theUser.SessionID2 == "" {
		session += "2"
	} else if theUser.SessionID3 == "" {
		session += "3"
	}

	var res model.User
	// Define opt to return the updated document
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter = bson.M{"email": email, "password": Hash(password)}
	update := bson.M{"$set": bson.M{session: Hash(sessionid)}}
	err = db.GetCollection().FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(&res)
	if err != nil {
		return model.User{}, err, ""
	}

	return res, nil, sessionid
}
