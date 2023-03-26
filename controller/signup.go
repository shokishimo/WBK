package controller

import (
	"context"
	"github.com/shokishimo/WhatsTheBestKeyboard/db"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet { // handle GET method
		RenderPage(w, "static/public/signup.html")
	} else if r.Method == http.MethodPost { // handle POST method
		errorMessage := signUpPost(w, r)
		if errorMessage != "" {
			// render error message
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		// if sign up ok, now passcode check
		// Redirect to account home page
		http.Redirect(w, r, "/verifyPasscode", http.StatusSeeOther)
	} else { // others
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

// signUpPost saves a user signed up
func signUpPost(w http.ResponseWriter, r *http.Request) string {
	email := r.FormValue("email")
	password := r.FormValue("password")
	result := ValidateSignupInput(email, password)
	if result != "" {
		w.WriteHeader(http.StatusBadRequest)
		return result
	}
	// parse username from email
	separatedEmail := strings.Split(email, "@")
	username := separatedEmail[len(separatedEmail)-2]
	passcode := GeneratePasscode()

	theUser := model.User{
		Username:  username,
		Email:     email,
		Password:  Hash(password),
		SessionID: passcode,
		Fav:       []model.Keyboard{},
		BestKeys:  []model.Keyboard{},
		WorstKeys: []model.Keyboard{},
	}

	client := db.Connect()
	collection := db.GetAccessKeysToUsersCollection(client)
	defer db.Disconnect(client)

	// check if the input user already exists in the database
	// Define the filter to find a specific document
	var res model.User
	filter := bson.M{"email": email}
	err := collection.FindOne(context.TODO(), filter).Decode(&res)
	// when the user with the sessionID found
	if err == nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return "http.StatusNotAcceptable; already a user with the same email exists"
	}

	// save the user temporary
	collection = db.GetAccessKeysToTemporaryUsersCollection(client)
	err = theUser.SaveUser(collection)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return "Failed to save the user"
	}

	// send email to let them validate their email address
	err = SendPasscodeMail(email, passcode)
	if err != nil {
		return "Failed to send mail to the user: " + err.Error()
	}

	// set email to the browser so that a user can let the system to resend their passcode in case there is a issue
	SetEmailCookie(w, email)
	return ""
}
