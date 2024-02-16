package main

import (
	// "fmt"
	"bufio"
	"fmt"
	"log"
	"os"
	// "github.com/gorilla/websocket"
)

var url string = "http://localhost:8000"


func main() {
		message := get_input()
    log.Println("server:", url)
		log.Println(message)
}

func get_input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Message:")
	message, _ := reader.ReadString('\n')
	return message
}