package server

import (
	"encoding/json"
	"fmt"
	"getsome-db/internal/db"
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
	// Initialize the database
	dbInstance, err := db.NewDatabase("storage/data.json")
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// Configure the HTTP server with the desired settings
	myServer := &http.Server{
		Addr:         port,                       // Port to listen on
		Handler:      setUpDbRoutes(dbInstance),  // Function to set up the HTTP routes
		ReadTimeout:  readTimeOut * time.Second,  // Read timeout to limit the maximum time for reading a request
		WriteTimeout: writeTimeOut * time.Second, // Write timeout to limit the maximum time for writing a response
		IdleTimeout:  idleTimeOut * time.Second,  // Idle timeout to close idle connections
	}

	log.Printf("Server is running on port: %s", port)

	// Start the HTTP server and begin listening for requests
	return myServer.ListenAndServe() // This will block and keep the server running
}

// setUpDbRoutes configures the HTTP routes for the database
func setUpDbRoutes(dbInstance *db.Database) http.Handler {
	// Create a new ServeMux (HTTP request multiplexer)
	mux := http.NewServeMux()

	// Define the health check route
	mux.HandleFunc("/health", healthCheckHandler) // Health check endpoint

	// CRUD routes with closures to use dbInstance
	mux.HandleFunc("/db/create", func(w http.ResponseWriter, r *http.Request) { createHandler(w, r, dbInstance) })
	mux.HandleFunc("/db/read", func(w http.ResponseWriter, r *http.Request) { readHandler(w, r, dbInstance) })
	mux.HandleFunc("/db/update", func(w http.ResponseWriter, r *http.Request) { updateHandler(w, r, dbInstance) })
	mux.HandleFunc("/db/delete", func(w http.ResponseWriter, r *http.Request) { deleteHandler(w, r, dbInstance) })

	return mux
}

// healthCheck is the handler function for the "/health" route
func healthCheckHandler(resW http.ResponseWriter, req *http.Request) {
	// Respond with a simple message to indicate the server is healthy
	fmt.Fprintf(resW, "Server is healthy!")
}

// create data handler
func createHandler(w http.ResponseWriter, r *http.Request, dbInstance *db.Database) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqData struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = dbInstance.Create(reqData.Key, reqData.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Data created successfully")
}

// read data handler
func readHandler(w http.ResponseWriter, r *http.Request, dbInstance *db.Database) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is required", http.StatusBadRequest)
		return
	}

	value, err := dbInstance.Read(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"key": key, "value": value})
}

// update data handler
func updateHandler(w http.ResponseWriter, r *http.Request, dbInstance *db.Database) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqData struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = dbInstance.Update(reqData.Key, reqData.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Data updated successfully")
}

// delete data handler
func deleteHandler(w http.ResponseWriter, r *http.Request, dbInstance *db.Database) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is required", http.StatusBadRequest)
		return
	}

	err := dbInstance.Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Data deleted successfully")
}
