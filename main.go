package main

import (
	"fmt"
	"learning_datastar/examples/counter"
	"net/http"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle the routes
	http.HandleFunc("/counter", counter.PageHandler)
	http.HandleFunc("/counter/increment", counter.CounterIncrementPutHandler)

	// Start the server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
