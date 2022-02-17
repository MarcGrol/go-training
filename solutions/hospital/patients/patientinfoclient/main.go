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
	patientUid := os.Args[1]
	client, cleanup, err := patientinfoapi.NewGrpcClient(patientinfoapi.DefaultPort)
	if err != nil {
		log.Fatalf("*** Error creating patient-info-client: %v", err)
	}
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetPatientOnUid(ctx, &patientinfoapi.GetPatientOnUidRequest{
		PatientUid: patientUid,
	})
	if err != nil {
		log.Fatalf("Error fetching client: %s", err)
	}

	log.Printf("Patient: %+v", resp.Patient)
}
