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
		router := router{handlers: map[string] map[string] func([]string){}}
		router.configureGET( "/", doGet)
		router.configurePOST("/", doPost)
		
		doStuff(router, input)
	}
}

func (r router) deferRequest(method string, path []string) {
	paths, ok := r.handlers[method]
	if !ok {
		fmt.Println("unknown method")
		return
	}
	handler, ok := paths[path[0]]
	if !ok {
		fmt.Println("404 - Path not found.")
		return
	}
	handler(path)

}

// Everything below is the "framework" 
type router struct {
	handlers map[string] map[string] func([]string)
}

func (r router) configureGET(path string, f func([]string)) {
	handler, ok := r.handlers["GET"]
	if !ok {
		handler = map[string] func([]string){}
		r.handlers["GET"] = handler
	}
	r.handlers["GET"][path] = f
}

func (r router) configurePOST(path string, f func([]string)) {
	handler, ok := r.handlers["POST"]
	if !ok {
		handler = map[string] func([]string){}
		r.handlers["POST"] = handler
	}
	r.handlers["POST"][path] = f
}

func doStuff(routes router, line string) {
	tokens := strings.Split(line, " ")
	routes.deferRequest(tokens[0], tokens[1:])

}
