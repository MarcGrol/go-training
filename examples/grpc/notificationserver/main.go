package main

import (
	"log"
	"time"

	_ "github.com/MarcGrol/go-training/examples/grpc/notificationapi"
)

const (
	port = ":50051"
)

func main() {
	s := New()
	time.Sleep(time.Second * 2)
	err := s.GRPCListenBlocking(port)
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
