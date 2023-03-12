package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func ForgetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("static/public/forget.html")
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
