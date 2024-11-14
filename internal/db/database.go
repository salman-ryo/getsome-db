package db

import (
	"encoding/json" // Provides functions to work with JSON encoding and decoding
	"errors"        // Standard package for error handling
	"fmt"
	"os"   // Provides functions to work with the operating system (e.g., checking if a file exists)
	"sync" // Provides concurrency control primitives, like mutexes
)

// Database struct defines the structure of the in-memory JSON database.
type Database struct {
	data     map[string]interface{} // In-memory data store, using a map to store key-value pairs
	filePath string                 // Path to the JSON file where data is stored persistently
	mu       sync.RWMutex           // Mutex for safe concurrent access; ensures only one goroutine modifies data at a time
}

// NewDatabase initializes the database and loads any existing data from a file.
func NewDatabase(databaseName string) (*Database, error) {
	filePath := "storage/" + databaseName + ".json"
	myDb := &Database{
		data:     make(map[string]interface{}), // Initializes the in-memory data store as an empty map
		filePath: filePath,                     // Sets the file path for persistence
	}

	if err := myDb.loadFromDisk(); err != nil { // Calls loadFromDisk and returns error if it fails
		return nil, fmt.Errorf("error loading database %s: %w", databaseName, err)
	}
	return myDb, nil
}

// loadFromDisk reads JSON data from the file at filePath and populates the database.
func (db *Database) loadFromDisk() error { // Declares loadFromDisk as a method of Database type
	db.mu.Lock()         // Acquires a write lock to prevent other goroutines from reading or writing
	defer db.mu.Unlock() // Releases the lock once function completes, ensuring safe concurrent access

	file, err := os.ReadFile(db.filePath) // Reads the entire file into memory
	if os.IsNotExist(err) {               // If the file doesn’t exist, start with an empty database and return nil
		return nil
	} else if err != nil { // If another error occurs (e.g., permission error), return it
		return err
	}

	// Unmarshal the JSON data from the file into the database's in-memory data store
	return json.Unmarshal(file, &db.data)
}

// saveToDisk writes the current state of the in-memory data store to the JSON file.
func (db *Database) saveToDisk() error {
	db.mu.Lock()         // Locks the data for writing, preventing concurrent access
	defer db.mu.Unlock() // Unlocks the mutex once data is written

	data, err := json.MarshalIndent(db.data, "", "  ") // Converts data to JSON format with indentation
	if err != nil {                                    // If an error occurs during marshalling, return it
		return err
	}

	// Writes the JSON data to the file with permissions set to 0644 (owner read/write, others read)
	return os.WriteFile(db.filePath, data, 0644) // Writes data to file with permissions 0644
}

// CRUD Functions:

// Create adds a new key-value pair to the database.
func (db *Database) Create(key string, value interface{}) error {
	db.mu.Lock()         // Locks for write access to prevent data races
	defer db.mu.Unlock() // Unlocks after function completes

	if _, exists := db.data[key]; exists { // Checks if key already exists
		return fmt.Errorf("Create error: key %s already exists", key) // Returns error if key exists
	}

	db.data[key] = value
	if err := db.saveToDisk(); err != nil {
		return fmt.Errorf("Create error: %w", err)
	}
	return nil // Saves updated data to disk for persistence
}

// Read retrieves the value associated with a key.
func (db *Database) Read(key string) (interface{}, error) {
	db.mu.RLock()         // Locks for reading to prevent data races
	defer db.mu.RUnlock() // Unlocks after function completes

	value, exists := db.data[key] // Looks up the value associated with the key
	if !exists {                  // If the key doesn’t exist, return an error
		return nil, errors.New("Read error: key not found")
	}
	return value, nil // Returns the value if key is found
}

// Update modifies the value associated with an existing key.
func (db *Database) Update(key string, value interface{}) error {
	db.mu.Lock()         // Locks for write access to prevent data races
	defer db.mu.Unlock() // Unlocks after function completes

	if _, exists := db.data[key]; !exists { // Checks if key exists
		return errors.New("Update error: key not found") // Returns error if key doesn't exist
	}

	db.data[key] = value // Updates the value for the specified key
	if err := db.saveToDisk(); err != nil {
		return fmt.Errorf("Update error: %w", err)
	}
	return nil
}

// Delete removes a key-value pair from the database.
func (db *Database) Delete(key string) error {
	db.mu.Lock()         // Locks for write access to prevent data races
	defer db.mu.Unlock() // Unlocks after function completes

	if _, exists := db.data[key]; !exists { // Checks if key exists in the data
		return errors.New("Delete error: key not found") // Returns error if key doesn't exist
	}

	delete(db.data, key) // Removes the key-value pair from the in-memory data store
	if err := db.saveToDisk(); err != nil {
		return fmt.Errorf("Delete error: %w", err)
	}
	return nil
}
