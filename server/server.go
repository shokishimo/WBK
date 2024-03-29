package server

import (
	"github.com/shokishimo/WhatsTheBestKeyboard/controller"
	"net/http"
)

// ServeMux creates a new HTTP server.
func ServeMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.HomeHandler)
	mux.HandleFunc("/signup", controller.SignUpHandler)
	mux.HandleFunc("/verifyPasscode", controller.PasscodeVerificationHandler)
	mux.HandleFunc("/login", controller.LoginHandler)
	mux.HandleFunc("/logout", controller.LogoutHandler)
	mux.HandleFunc("/search", controller.SearchHandler)
	mux.HandleFunc("/keyboarddetail", controller.KeyboardDetailHandler)
	mux.HandleFunc("/forget", controller.ForgetHandler)
	mux.HandleFunc("/createNewKeyboard", controller.CreateNewKeyboardHandler)
	mux.HandleFunc("/getRanking", controller.GetRankingHandler)
	mux.HandleFunc("/account", controller.AccountHandler)
	mux.HandleFunc("/newKeyboardRequest", controller.NewKeyboardRequestHandler)
	mux.HandleFunc("/notFound", controller.Handle404)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	return mux
}
