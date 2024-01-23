package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! Go is here!")
	}

	// Register the handler for the root route "/"
	http.HandleFunc("/", handler)

	// Specify the port to listen on
	port := 7777

	// Print a message indicating the server is running
	fmt.Printf("Server is running on :%d...\n", port)

	// Start the HTTP server on the specified port
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
