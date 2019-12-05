package main

import (
	"log"

	"github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"
)

func main() {
	s := newServer(newPatientStore())
	err := s.GRPCListenBlocking(patientinfoapi.DefaultPort)
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
