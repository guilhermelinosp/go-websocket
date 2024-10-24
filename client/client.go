package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost:6372/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error connecting to the WebSocket server:", err)
	}
	defer conn.Close()

	log.Println("Connected to the server")

	for {
		msg := "Hello, WebSocket!"
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}

		log.Printf("Sent: %s\n", msg)

		_, response, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received: %s\n", response)

		time.Sleep(2 * time.Second)
	}
}
