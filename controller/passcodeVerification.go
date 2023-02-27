package controller

import (
	"html/template"
	"net/http"
	"strconv"
)

func PasscodeVerificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		VerifyPassGet(w)
	} else if r.Method == http.MethodPost {
		VerifyPassPost(w, r)
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
func VerifyPassPost(w http.ResponseWriter, r *http.Request) {
	var inPasscode string
	for i := 1; i <= 6; i++ {
		inPasscode = inPasscode + r.FormValue("in"+strconv.Itoa(i))
	}
	w.Write([]byte(inPasscode))

	// Once passcode is verified, create a sessionID and cookie and set them up for the user
	//// save the cookie in the client browser
	//user.SetCookie(w, sessionID)
	//

}
