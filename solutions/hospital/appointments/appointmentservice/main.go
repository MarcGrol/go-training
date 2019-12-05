package main

import (
	"log"

	"github.com/google/uuid"

	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentapi"
)

type uuider struct{}

func (u uuider) Create() string {
	return uuid.New().String()
}
func main() {
	s := newServer(newAppointmentStore(uuider{}))
	err := s.GRPCListenBlocking(appointmentapi.DefaultPort)
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
