package controller

import (
	"fmt"
	"net/http"
)

func ForgetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		RenderPage(w, "static/public/forget.html")
	} else if r.Method == http.MethodPost {
		err := handleForgetPost(w, r)
		if err != "" {
			http.Error(w, err, http.StatusInternalServerError)
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	return
}

func handleForgetPost(w http.ResponseWriter, r *http.Request) string {
	email := r.FormValue("email")
	if ValidateEmail(email) != "" {
		return "email is invalid"
	}

	// When email is valid
	// TODO: Implement the rest
	fmt.Println(email)

	return ""
}
