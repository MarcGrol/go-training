
package main

import (
	_ "github.com/MarcGrol/go-training/examples/grpc/spec"
	"log"
)

func main() {
	err := New().ListenHttpBlocking(":50051", ":8080")
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}

