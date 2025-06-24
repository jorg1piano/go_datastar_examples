package main

import (
	"fmt"
	"learning_datastar/examples/count_until"
	"learning_datastar/examples/counter"
	"learning_datastar/examples/dialog"
	"net/http"
)

func main() {
	// Handle the routes
	http.HandleFunc("/counter", counter.PageHandler)
	http.HandleFunc("/counter/increment", counter.CounterIncrementPutHandler)

	http.HandleFunc("/count-until", count_until.PageHandler)
	http.HandleFunc("/count-until/increment-slowly", count_until.PutHandler)

	http.HandleFunc("/dialog", dialog.PageHandler)
	http.HandleFunc("/dialog/quote", dialog.GetHandler)

	// Start the server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
