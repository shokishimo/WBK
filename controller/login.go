package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderLoginPage(w)
	} else if r.Method == http.MethodPost {

	}
}

func renderLoginPage(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("static/public/login.html")
	if err != nil {
		fmt.Println("Error: Failed to get login page")
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error: Failed to render login page")
		return
	}
}
