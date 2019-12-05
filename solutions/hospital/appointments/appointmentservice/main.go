package main

import (
	"log"

	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"

	"github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"

	"github.com/google/uuid"

	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentapi"
)

type uuider struct{}

func (u uuider) Create() string {
	return uuid.New().String()
}

func main() {
	log.Printf("Startup")

	patientClient, patientClientCleanup, err := patientinfoapi.NewGrpcClient(patientinfoapi.DefaultPort)
	if err != nil {
		log.Fatalf("*** Error creating patient-info-client: %v", err)
	}
	defer patientClientCleanup()
	log.Printf("Created patient-client")

	client, cleanup, err := notificationapi.NewGrpcClient(notificationapi.DefaultPort)
	if err != nil {
		log.Fatalf("*** Error creating motification-client: %v", err)
	}
	defer cleanup()

	log.Printf("Created notification-client")

	log.Printf("Starting appointment server")
	s := newServer(newAppointmentStore(uuider{}), patientClient, client)
	err = s.GRPCListenBlocking(appointmentapi.DefaultPort)
	if err != nil {
		log.Fatalf("Error starting appointmenmt server: %s", err)
	}
}
