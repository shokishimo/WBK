package controller

import (
	"net/http"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	RenderPage(w, "static/public/account.html")
}
