package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// These are our handler functions.
func doGet(what Context) {
	what.String(200, "GET: %v\n", what.FullPath())
}

func doPost(what Context) {
	what.String(200 , "POST: %v\n", what.FullPath())
}

// This is the "main" where we configure our framework to handle the CLI commands
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// the router is what directs the get/post to the handlers
		router := router{handlers: map[string] map[string] func(Context){}}
		router.configureGET( "/", doGet)
		router.configurePOST("/", doPost)
		
		doStuff(router, input)
	}
}
// Everything below is the "framework" 

func (r router) deferRequest(method string, path []string) {
	context := Context{
		fullPath : "",
		request : Request{ method: method},
	}
	// handle error input where there is no path
	if len (path)  > 0 {
		context.fullPath = path[0]
	}

	paths, ok := r.handlers[method]
	if !ok {
		context.String( 405, "Method not allowed: \n", method)
		return
	}
	handler, ok := paths[path[0]]
	if !ok {
		context.String( 404, "Path not found: \n", path[0])
		return
	}

	handler(context)
}

type router struct {
	handlers map[string] map[string] func(Context)
}

type Context struct {
	request Request
	fullPath string
}

func (c Context) FullPath() string {
	return c.fullPath
}

func (c Context) String(code int, format string, values ...interface{}) {
	fmt.Printf("Return code %v, %v\n", code, values)
}

type Request struct {
	method string
}

func (r router) configureGET(path string, f func(Context)) {
	handler, ok := r.handlers["GET"]
	if !ok {
		handler = map[string] func(Context){}
		r.handlers["GET"] = handler
	}
	r.handlers["GET"][path] = f
}

func (r router) configurePOST(path string, f func(Context)) {
	handler, ok := r.handlers["POST"]
	if !ok {
		handler = map[string] func(Context){}
		r.handlers["POST"] = handler
	}
	r.handlers["POST"][path] = f
}

func doStuff(routes router, line string) {
	tokens := strings.Split(line, " ")
	routes.deferRequest(tokens[0], tokens[1:])

}
