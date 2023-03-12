package controller

import (
	"html/template"
	"net/http"
)

func ForgetHandler(w http.ResponseWriter, r *http.Request) {
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
}
