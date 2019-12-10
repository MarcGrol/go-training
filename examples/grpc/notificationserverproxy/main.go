package main

import (
	"log"

	_ "github.com/MarcGrol/go-training/examples/grpc/notificationapi"
)

const (
	port = ":8080"
)

func main() {
	err := New().ListenHttpBlocking(":50051", port)
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
