package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// These are our handler functions.
func doGet(what []string) {
	fmt.Printf("GET: %v\n", what)
}

func doPost(what []string) {
	fmt.Printf("POST: %v\n", what)
}

// This is the "main" where we configure our framework to handle the CLI commands
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// the router is what directs the get/post to the handlers
		router := router{handlers: map[string] func([]string){}}
		router.configureGET(doGet)
		router.configurePOST(doPost)
		
		doStuff(router, input)
	}
}

// Everything below is the "framework" 
type router struct {
	handlers map[string] func([]string)
}

func (r router) configureGET(f func([]string)) {
	r.handlers["GET"] = f
}

func (r router) configurePOST(f func([]string)) {
	r.handlers["POST"] = f
}

func doStuff(routes router, line string) {
	tokens := strings.Split(line, " ")

	switch tokens[0] {
	case "GET":
		routes.handlers["GET"](tokens[1:])
	case "POST":
		routes.handlers["POST"](tokens[1:])
	default:
		fmt.Printf("unknown command\n")
	}
}
