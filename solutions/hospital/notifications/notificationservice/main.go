package main

import (
	"log"
)

func main() {
	s := newServer()
	err := s.GRPCListenBlocking(":60003")
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
