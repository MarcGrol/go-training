package main

import (
	"log"

	"github.com/google/uuid"
)

type uuider struct{}

func (u uuider) Create() string {
	return uuid.New().String()
}
func main() {
	s := newServer(newAppointmentStore(uuider{}))
	err := s.GRPCListenBlocking(":60001")
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
