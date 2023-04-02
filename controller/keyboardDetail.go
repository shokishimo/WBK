package controller

import (
	"fmt"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"html/template"
	"net/http"
)

func KeyboardDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/keyboarddetail" {
		http.Redirect(w, r, "/notFound", http.StatusTemporaryRedirect)
	}
	if r.Method == http.MethodGet {
		if err := keyboardDetail(w, r); err != nil {
			http.Error(w, fmt.Sprintf("Error: Handling the request: %s", err.Error()), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "HTTP Request Method NOT allowed", http.StatusMethodNotAllowed)
	}
}

func keyboardDetail(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	name := r.FormValue("name")
	ranking := r.FormValue("ranking")

	err, theKeyboard := model.FindKeyboardWithNameAndRanking(name, ranking)
	if err != nil {
		return err
	}

	// render a page
	tmpl, err := template.ParseFiles("static/public/keyboardDetail.html")
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return nil
	}

	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, theKeyboard)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return nil
}
