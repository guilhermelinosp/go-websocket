package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()

	log.Printf("Client connected: %s\n", conn.RemoteAddr())

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error while reading message:", err)
			break
		}

		log.Printf("Received: %s\n", msg)

		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			log.Println("Error while writing message:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)

	log.Println("Server is running on :6372")
	log.Fatal(http.ListenAndServe(":6372", nil))
}
