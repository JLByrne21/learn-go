package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Client")
}


func get_input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Message:")
	message, _ := reader.ReadString('\n')
	return message
}
