package main

import (
	"log"
	"time"
)

const (
	port = ":50051"
)

func main() {
	time.Sleep(time.Second * 2)
	_, err := New(port)
	if err != nil {
		log.Fatalf("Error starting rest-notification service: %s", err)
	}
}
