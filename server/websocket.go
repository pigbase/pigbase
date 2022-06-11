package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Loop(conn *websocket.Conn) {
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: Message: %s, Message Type: %d", message, messageType)
		req := Request{}
		err = json.Unmarshal(message, &req)
		fmt.Printf("%v", req)
	}
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	Loop(conn)
}

func InitWebsocket() {
	log.Println("Starting websocket server")
	http.HandleFunc("/socket", SocketHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
