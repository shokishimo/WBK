package controller

import (
	"context"
	"errors"
	"github.com/shokishimo/WhatsTheBestKeyboard/database"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	RenderPageWithStringData(w, "static/public/login.html", data.ErrorString)
}

func handleLogin(w http.ResponseWriter, r *http.Request) string {
	email := r.FormValue("email")
	password := r.FormValue("password")
	result := ValidateSignupInput(email, password)
	if result != "" {
		w.WriteHeader(http.StatusBadRequest)
		return result
	}

	err := LoginSessions(w, email, password)
	if err != nil {
		return err.Error()
	}

	return ""
}

func LoginSessions(w http.ResponseWriter, email string, password string) error {
	db := database.Connect()
	defer db.Disconnect()
	db.GetAccessKeysToUsersCollection()

	var theUser model.User
	filter := bson.M{"email": email, "password": Hash(password)}
	err := db.GetCollection().FindOne(context.TODO(), filter).Decode(&theUser)
	// when the user with the passcode not found
	if err != nil {
		return err
	}

	sessionid := GenerateSessionID()
	var sessionNum string
	// When 3 devices are filled
	if theUser.SessionID1 != "" && theUser.SessionID2 != "" && theUser.SessionID3 != "" {
		return errors.New("one can access their account with up to 3 devices. No more device is available")
	} else if theUser.SessionID1 == "" {
		sessionNum = "1"
	} else if theUser.SessionID2 == "" {
		sessionNum = "2"
	} else if theUser.SessionID3 == "" {
		sessionNum = "3"
	}
	session := "sessionid" + sessionNum

	var res model.User
	// Define opt to return the updated document
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter = bson.M{"email": email, "password": Hash(password)}
	update := bson.M{"$set": bson.M{session: Hash(sessionid)}}
	err = db.GetCollection().FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(&res)
	if err != nil {
		return err
	}

	SetUsernameCookie(w, theUser.Username)
	SetSessionCookie(w, sessionNum, sessionid)
	SetSessionNumInCookie(w, sessionNum)

	return nil
}
