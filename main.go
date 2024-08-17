package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"net/http"
    "github.com/gin-gonic/gin"
	"bytes"
)

// These are our handler functions.
func doGet(c *gin.Context) {
	c.String(200, "GET: %v\n", c.Request.URL.Path)
}

func doPost(c *gin.Context) {
	c.String(200 , "POST: %v\n", c.Request.URL.Path)
}

// This is the "main" where we configure our framework to handle the CLI commands
func main() {
	reader := bufio.NewReader(os.Stdin)
	
	// the router is what directs the get/post to the handlers
	router := router{handlers: map[string] map[string] func(*gin.Context){}}
	router.configureGET( "/", doGet)
	router.configurePOST("/", doPost)

	gin.SetMode(gin.ReleaseMode)
	ginRouter := gin.Default()
	ginRouter.GET("/", doGet)
	ginRouter.POST("/", doPost)
	
	go func() { ginRouter.Run("localhost:8080") }()
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		doStuff(router, input)
	}
}

// Everything below is the "framework" 
func doStuff(routes router, line string) {
	tokens := strings.Split(line, " ")
	if len(tokens) < 2 {
		fmt.Printf("error: please provide at least 2 tokens on the CLI.\n")
		return
	}
	routes.deferRequest(tokens[0], tokens[1:])
}

func (r router) deferRequest(method string, path []string) {
	// gin will have an internal way of constructing this.
	c,_ := gin.CreateTestContext(CLIResponseWriter{})

	paths, ok := r.handlers[method]
	if !ok {
		c.String( 405, "Method not allowed: %v \n", method)
		return
	}
	handler, ok := paths[path[0]]
	if !ok {
		c.String( 404, "Path not found: %v \n", path[0])
		return
	}
	c.Request, _ = http.NewRequest(http.MethodPost, path[0], bytes.NewBuffer([]byte("{}")))
	handler(c)
}

type router struct {
	handlers map[string] map[string] func(*gin.Context)
}

// A couple of functions to impersonate the http writer
type CLIResponseWriter struct {

}

func (CLIResponseWriter) Header() http.Header {
	return map[string][]string{}
}
func (CLIResponseWriter) Write(b []byte) (int, error)  {
	written, error := os.Stdout.Write(b)
	// fmt.Println("Write")
	return written, error
}
func (CLIResponseWriter) WriteHeader(statusCode int) {
	fmt.Printf("WriteHeader: %v\n ",statusCode)
}

// Back to "our" framework code
func (r router) configureGET(path string, f func(*gin.Context)) {
	handler, ok := r.handlers["GET"]
	if !ok {
		handler = map[string] func(*gin.Context){}
		r.handlers["GET"] = handler
	}
	r.handlers["GET"][path] = f
}

func (r router) configurePOST(path string, f func(*gin.Context)) {
	handler, ok := r.handlers["POST"]
	if !ok {
		handler = map[string] func(*gin.Context){}
		r.handlers["POST"] = handler
	}
	r.handlers["POST"][path] = f
}
