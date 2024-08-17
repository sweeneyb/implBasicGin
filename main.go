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

type router struct {
	handlers map[string] func([]string)
}

func doStuff(line string) {
	tokens := strings.Split(line, " ")

	router := router{handlers: map[string] func([]string){}}
	router.handlers["GET"] = doGet
	router.handlers["POST"] = doPost
	
	switch tokens[0] {
	case "GET":
		router.handlers["GET"](tokens[1:])
	case "POST":
		router.handlers["POST"](tokens[1:])
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
