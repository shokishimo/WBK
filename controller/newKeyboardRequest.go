package controller

import "net/http"

func NewKeyboardRequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		RenderPage(w, "static/public/newKeyboardRequest.html")
	}
}
