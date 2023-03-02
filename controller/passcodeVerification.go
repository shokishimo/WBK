package controller

import (
	"context"
	"fmt"
	"github.com/shokishimo/WhatsTheBestKeyboard/db"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"go.mongodb.org/mongo-driver/bson"
	"html/template"
	"net/http"
	"strconv"
)

func PasscodeVerificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		VerifyPassGet(w)
	} else if r.Method == http.MethodPost {
		err := VerifyPassPost(w, r)
		if err != nil {
			w.WriteHeader(http.StatusNotAcceptable)
			_, err2 := w.Write([]byte(err.Error()))
			if err2 != nil {
				return
			}
			return
		}
		// if status accepted
		// Redirect to account home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte("http method allowed"))
		if err != nil {
			return
		}
		return
	}
}

// VerifyPassGet handles HTTP GET request
func VerifyPassGet(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("static/public/verifyPasscode.html")
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// VerifyPassPost handles HTTP Post request
func VerifyPassPost(w http.ResponseWriter, r *http.Request) error {
	var inPasscode string
	for i := 1; i <= 6; i++ {
		inPasscode = inPasscode + r.FormValue("in"+strconv.Itoa(i))
	}

	// validate if the passcode is correct
	client := db.Connect()
	defer db.Disconnect(client)
	collection := db.GetAccessKeysToTemporaryUsersCollection(client)

	// check if the input user already exists in the database
	// Define the filter to find a specific document
	var theUser model.User
	filter := bson.M{"sessionid": inPasscode}
	err := collection.FindOne(context.TODO(), filter).Decode(&theUser)
	// when the user with the passcode not found
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Once passcode is verified, create and set session id
	sessionId := GenerateSessionID()
	theUser.SessionID = Hash(sessionId)

	// save the sessionid and username in the client browser
	SetSessionCookie(w, sessionId)
	SetUsernameCookie(w, theUser.Username)
	// delete email cookie
	DeleteCookie(w, "email", theUser.Email)

	// delete this user from the temporary and save user to the users table
	err = model.DeleteUser(theUser, collection)
	collection = db.GetAccessKeysToUsersCollection(client)
	err = model.SaveUser(theUser, collection)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
