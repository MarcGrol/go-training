package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s patient-uid", os.Args[0])
	}
	client, cleanup, err := patientinfoapi.NewGrpcClient(patientinfoapi.DefaultPort)
	if err != nil {
		log.Fatalf("*** Error creating patient-info-client: %v", err)
	}
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetPatientOnUid(ctx, &patientinfoapi.GetPatientOnUidRequest{
		PatientUid: os.Args[1],
	})
	if err != nil {
		log.Fatalf("Error fetching client: %s", err)
	}
	if resp.Error != nil {
		log.Fatalf("Error fetching client: %+v", resp.Error)
	}
	log.Printf("Patient: %+v", resp.Patient)
}
