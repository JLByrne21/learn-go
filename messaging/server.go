package main

import (
	// "fmt"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var url string = "http://localhost:8000"
var upgrader = websocket.Upgrader{}
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func main() {
	log.Println("server:", url)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Failed to start server")
		return
	}

}


func handleConnections(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
			log.Println(err)
			return
	}
	defer conn.Close()
	clients[conn] = true

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
		 return
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		message := <- broadcast
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my GOLang chat app!")
}