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
	} else if r.Method == http.MethodPost {

	}
}
