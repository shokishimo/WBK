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
	path := r.URL.Path
	if path != "/" { // if endpoint is unknown
		// Redirect to 404 page
		http.Redirect(w, r, "/notFound?path="+path, http.StatusSeeOther)
	}

	ServePublicHome(w, r)
}

// ServePublicHome shows the public template home to the browser
func ServePublicHome(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/public/home.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// get keyboards ranking
	Keyboards := model.GetRanks(4)
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, Keyboards)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
