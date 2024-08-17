package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		tokens := strings.Split(input, " ")
		switch tokens[0] {
		case "GET":
			fmt.Printf("GET: %v\n", tokens[1:])
		case "GET":
			fmt.Printf("GET: %v\n", tokens[1:])
		default:
			fmt.Printf("unknown command\n")
		}
	}
}
