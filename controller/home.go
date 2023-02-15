package controller

import (
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ServePublicHome(w)
}

// ServePublicHome shows the public template home to the browser
func ServePublicHome(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("static/public/home.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	keyboards := model.GetRanks(4)
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, keyboards)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
