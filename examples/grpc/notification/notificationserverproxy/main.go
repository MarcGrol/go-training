package main

import (
	"log"

	_ "github.com/MarcGrol/go-training/examples/grpc/notification/notificationapi"
)

func main() {
	err := New().ListenHttpBlocking(":50051", ":8082")
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
