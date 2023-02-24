package main

import (
	"fmt"
	"github.com/shokishimo/WhatsTheBestKeyboard/controller"
	"github.com/shokishimo/WhatsTheBestKeyboard/server"
	"log"
	"net/http"
)

func main() {
	err := controller.SendPasscodeMail("soutaw0429@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	srvMux := server.ServeMux()
	fmt.Println("Server is running")
	if err := http.ListenAndServe(":3000", srvMux); err != nil {
		log.Fatal("Error in running the server")
		fmt.Printf("e: %v\n", err)
	}
}
