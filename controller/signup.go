package controller

import (
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
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

	theUser := model.CreatNewUser(username, email, Hash(password))
	passcode := GeneratePasscode()
	theUser.SessionID1 = passcode

	// check if a user is already in the database
	_, err := model.FindUserWithEmail(email)
	if err == nil {
		return "http.StatusNotAcceptable; already a user with the same email exists"
	}

	// save the user temporary
	err = theUser.SaveUserToTemporary()
	if err != nil {
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
