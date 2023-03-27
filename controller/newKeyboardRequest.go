package controller

import (
	"fmt"
	"net/http"
)

func NewKeyboardRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/newKeyboardRequest" {
		http.Redirect(w, r, "/notFound", http.StatusTemporaryRedirect)
	}
	if r.Method == http.MethodGet {
		RenderPage(w, "static/public/newKeyboardRequest.html")
	} else if r.Method == http.MethodPost {
		if err := handleNewKeyboardRequest(r); err != nil {
			http.Error(w, fmt.Sprintf("Error: Handling the request: %s", err.Error()), http.StatusBadRequest)
			return
		}
		RenderPage(w, "static/public/KBRequestSuccess.html")
	} else {
		http.Error(w, "HTTP Request Method NOT allowed", http.StatusMethodNotAllowed)
	}
}

func handleNewKeyboardRequest(r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	nickname := r.FormValue("nickname")
	keyboard := r.FormValue("keyboard")
	url := r.FormValue("url")
	if err := SendKBRequestMailToOwner(nickname, keyboard, url); err != nil {
		return err
	}

	return nil
}
