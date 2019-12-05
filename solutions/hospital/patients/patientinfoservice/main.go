package main

import (
	"log"
)

func main() {
	s := newServer(newPatientStore())
	err := s.GRPCListenBlocking(":60002")
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
