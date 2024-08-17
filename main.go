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
		doStuff(input)
	}
}

func doStuff(line string) {
	tokens := strings.Split(line, " ")
	switch tokens[0] {
	case "GET":
		doGet(tokens[1:])
	case "POST":
		doPost(tokens[1:])
	default:
		fmt.Printf("unknown command\n")
	}
}

func doGet(what []string) {
	fmt.Printf("GET: %v\n", what)
}

func doPost(what []string) {
	fmt.Printf("POST: %v\n", what)
}
