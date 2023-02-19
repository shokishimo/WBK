package controller

import (
	"html/template"
	"net/http"
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
		signUpPost(w, r)

	} else { // others
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
}

// signUpPost saves a user signed up
func signUpPost(w http.ResponseWriter, r *http.Request) {
	//sessionID := user.GenerateSessionID()
	//theUser := user.User{
	//	Username:  r.FormValue("username"),
	//	Password:  user.Hash(r.FormValue("password")),
	//	SessionID: user.Hash(sessionID),
	//}
	//// TODO: validate the user input
	//
	//// TODO: check if the input user already exists in the database
	//
	//// save the user
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
