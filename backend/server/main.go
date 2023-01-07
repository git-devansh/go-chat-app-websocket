package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/chat-app/server/handlers"
)

func setupRoutes() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/ws", handlers.WsEndpoint)
}

func main() {
	setupRoutes()

	// listening to chan
	go handlers.ListenToWsChannel()

	fmt.Println("-----\nServing at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
