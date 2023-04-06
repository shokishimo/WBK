package controller

import (
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
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
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		// if status accepted
		// Redirect to account home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("http method allowed"))
		return
	}
}

// VerifyPassGet handles HTTP GET request
func VerifyPassGet(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("static/public/verifyPasscode.html")
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
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
	inPasscode := ""
	for i := 1; i <= 6; i++ {
		inPasscode += r.FormValue("in" + strconv.Itoa(i))
	}

	// validate if the passcode is correct
	theUser, err := model.FindUserWithPasscode(inPasscode)
	if err != nil {
		return err
	}

	// Once passcode is verified, create and set session id
	sessionId1 := GenerateSessionID()
	theUser.SessionID1 = Hash(sessionId1)

	// save the sessionid and username in the client browser
	SetSessionCookie(w, sessionId1)
	SetUsernameCookie(w, theUser.Username)
	// delete email cookie
	DeleteCookie(w, "email", theUser.Email)

	// delete this user from the temporary and save user to the users table
	err = theUser.DeleteUserFromTemporary()
	err = theUser.SaveUser()
	if err != nil {
		return err
	}

	return nil
}
