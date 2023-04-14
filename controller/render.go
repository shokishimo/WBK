package controller

import (
	"html/template"
	"net/http"
)

func RenderPage(w http.ResponseWriter, path string) {
	tmpl, err := template.ParseFiles(path)
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

func RenderPageWithStringData(w http.ResponseWriter, path string, data string) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}
