package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	for {
		// messageType, message, err := conn.ReadMessage()
		// if err != nil {
		// 	log.Println("Error during message reading:", err)
		// 	break
		// }
		// log.Printf("Received: %s, %d", message, messageType)
		// err = conn.WriteMessage(messageType, message)
		// if err != nil {
		// 	log.Println("Error while sending message:", err)
		// 	break
		// }
	}
}

func Init() {
	log.Fatal("Starting websocket server")
	http.HandleFunc("/socket", socketHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
