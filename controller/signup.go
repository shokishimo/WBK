package controller

import (
	"context"
	"github.com/shokishimo/WhatsTheBestKeyboard/db"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"go.mongodb.org/mongo-driver/bson"
	"html/template"
	"net/http"
	"strings"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("static/public/signup.html")
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

	} else if r.Method == http.MethodPost { // handle POST method
		errorMessage := signUpPost(w, r)
		if errorMessage != "" {
			// render error message

		}
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
	if !result {
		w.WriteHeader(http.StatusBadRequest)
		return "http.StatusBadRequest"
	}
	sessionID := GenerateSessionID()
	separatedEmail := strings.Split(email, "@")
	username := separatedEmail[len(separatedEmail)-2]

	theUser := model.User{
		Username:  username,
		Email:     email,
		Password:  Hash(password),
		SessionID: Hash(sessionID),
		Fav:       []model.Keyboard{},
		BestKeys:  []model.Keyboard{},
		WorstKeys: []model.Keyboard{},
	}

	client := db.Connect()
	collection := db.GetAccessKeysToUsersCollection(client)
	defer db.Disconnect(client)

	// check if the input user already exists in the database
	// Define the filter to find a specific document
	filter := bson.M{"email": email}
	doesExists := collection.FindOne(context.TODO(), filter).Err()
	if doesExists != nil { // There is already a user with the email
		w.WriteHeader(http.StatusNotAcceptable)
		return "http.StatusNotAcceptable; already a user with the same email exists"
	}

	// TODO: send email to let them validate the their email address

	// save the user

	//err := user.SaveUser(theUser)
	//if err != nil {
	//	fmt.Fprint(w, err.Error())
	//	return
	//}
	//// success log
	//fmt.Println("successfully inserted the user")
	//
	//// save the cookie in the client browser
	//user.SetCookie(w, sessionID)
	//
	//// Redirect to account home page
	//http.Redirect(w, r, "/", http.StatusSeeOther)
}
