package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Constants for port and timeout values (in seconds)
const (
	port         = ":8080"
	readTimeOut  = 5   // Read timeout in seconds
	writeTimeOut = 10  // Write timeout in seconds
	idleTimeOut  = 120 // Idle timeout in seconds
)

// Start function to initialize and run the HTTP server
func Start() error {
	// Configure the HTTP server with the desired settings
	myServer := &http.Server{
		Addr:         port,                       // Port to listen on
		Handler:      setUpDbRoutes(),            // Function to set up the HTTP routes
		ReadTimeout:  readTimeOut * time.Second,  // Read timeout to limit the maximum time for reading a request
		WriteTimeout: writeTimeOut * time.Second, // Write timeout to limit the maximum time for writing a response
		IdleTimeout:  idleTimeOut * time.Second,  // Idle timeout to close idle connections
	}

	log.Printf("Server is running on port: %s", port)

	// Start the HTTP server and begin listening for requests
	return myServer.ListenAndServe() // This will block and keep the server running
}

// setUpDbRoutes configures the HTTP routes for the database
func setUpDbRoutes() http.Handler {
	// Create a new ServeMux (HTTP request multiplexer)
	mux := http.NewServeMux()

	// Define the health check route
	mux.HandleFunc("/health", healthCheck) // Health check endpoint

	return mux
}

// healthCheck is the handler function for the "/health" route
func healthCheck(resW http.ResponseWriter, req *http.Request) {
	// Respond with a simple message to indicate the server is healthy
	fmt.Fprintf(resW, "Server is healthy!")
}
