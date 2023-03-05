package controller

import (
	"html/template"
	"net/http"
)

type LoginData struct {
	ErrorString string
}

func renderLoginPage(w http.ResponseWriter, data LoginData) {
	tmpl, err := template.ParseFiles("static/public/login.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error: Failed to get login page"))
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error: Failed to render login page"))
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderLoginPage(w, LoginData{})
	} else if r.Method == http.MethodPost {
		res := handleLogin(w, r)
		if res != "" {
			renderLoginPage(w, LoginData{ErrorString: res})
		}

	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) string {
	email := r.FormValue("email")
	password := r.FormValue("password")
	result := ValidateSignupInput(email, password)
	if result != "" {
		w.WriteHeader(http.StatusBadRequest)
		return result
	}

	// check in database

	return ""
}
