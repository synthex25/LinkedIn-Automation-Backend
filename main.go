package main

import (
	"fmt"
	"net/http"
)

/*
main is the entry point of the application.

It initializes the HTTP server, registers all API routes,
attaches middleware for logging, and starts listening for
incoming requests.
*/
func main() {

	/*
		Register API routes.

		Each endpoint is wrapped with loggingMiddleware to
		provide basic observability for incoming requests.
	*/
	http.HandleFunc("/health", loggingMiddleware(healthHandler))
	http.HandleFunc("/send-message", loggingMiddleware(sendMessageHandler))
	http.HandleFunc("/send-connection", loggingMiddleware(sendConnectionHandler))

	// Log server startup information
	fmt.Println("Server running on http://localhost:9090")

	/*
		Start the HTTP server.

		ListenAndServe blocks the main goroutine and handles
		all incoming requests until the server is stopped
		or an unrecoverable error occurs.
	*/
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
