package controller

import (
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderLoginPage(w)
	} else if r.Method == http.MethodPost {
		err := handleLogin(w, r)
		if err != nil {

		}

	}
}

func renderLoginPage(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("static/public/login.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error: Failed to get login page"))
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error: Failed to render login page"))
		return
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) error {
	return nil
}