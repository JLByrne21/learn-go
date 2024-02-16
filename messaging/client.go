package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := get_input()
}


func get_input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Message:")
	message, _ := reader.ReadString('\n')
	return message
}
