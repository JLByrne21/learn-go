package main

import (
	// "fmt"

	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var url string = "http://localhost:8000"
var upgrader = websocket.Upgrader{}


func main() {
    log.Println("server:", url)
		http.HandleFunc("/send", handler)
		log.Fatal(http.ListenAndServe(":8000", nil))
}


func handler(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
			log.Println(err)
			return
	}
	defer conn.Close()
	for {
    messageType, p, err := conn.ReadMessage()
    if err != nil {
        log.Println(err)
        return
    }
    if err := conn.WriteMessage(messageType, p); err != nil {
        log.Println(err)
        return
    }
}
}
