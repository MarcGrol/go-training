package main

import (
	"log"

	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"

	"github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"

	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentapi"
)

func main() {
	log.Printf("Startup")

	patientClient, patientClientCleanup, err := patientinfoapi.NewGrpcClient(patientinfoapi.DefaultPort)
	if err != nil {
		log.Fatalf("*** Error creating patientinfo-client: %v", err)
	}
	defer patientClientCleanup()
	log.Printf("Created patientinfo-client")

	notificationClient, notificationClientCleanup, err := notificationapi.NewGrpcClient(notificationapi.DefaultPort)
	if err != nil {
		log.Fatalf("*** Error creating motification-client: %v", err)
	}
	defer notificationClientCleanup()

	log.Printf("Created notification-client")

	log.Printf("Starting appointment server")
	s := newServer(newAppointmentStore(NewBasicUuider()), patientClient, notificationClient)
	err = s.GRPCListenBlocking(appointmentapi.DefaultPort)
	if err != nil {
		log.Fatalf("Error starting appointment server: %s", err)
	}
}
