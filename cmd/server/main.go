package main

import (
	"fmt"
	"getsome-db/internal/server"
	"log"
)

func main() {
	fmt.Println("Starting GetSome server..")
	err := server.Start()

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
