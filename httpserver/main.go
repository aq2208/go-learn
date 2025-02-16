package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", nil)
}

// The main function is the entry point of the program. It is the first function that gets called when the program is executed. In this case, the main function sets up a simple HTTP server that listens on port 8080 and responds with "Hello, World!" when a request is made to the root path ("/").

// The http.HandleFunc function is used to register a handler function for a specific path. In this case, we are registering a handler function for the root path ("/"). The handler function takes two arguments: a http.ResponseWriter and a http.Request. The http.ResponseWriter is used to write the response back to the client, and the http.Request contains information about the incoming request.