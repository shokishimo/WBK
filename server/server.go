package server

import (
	"github.com/shokishimo/WhatsTheBestKeyboard/controller"
	"net/http"
)

// ServeMux creates a new HTTP server.
func ServeMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.HomeHandler)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	return mux
}
