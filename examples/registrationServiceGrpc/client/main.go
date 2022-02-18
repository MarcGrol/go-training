package main

import (
	"context"
	regprotobuf2 "github.com/MarcGrol/go-training/examples/registrationServiceGrpc/regprotobuf"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s bsn name", os.Args[0])
	}
	bsn := os.Args[1]
	name := os.Args[2]

	client, cleanup, err := regprotobuf2.NewGrpcClient(regprotobuf2.DefaultPort)
	if err != nil {
		log.Fatalf("*** Error creating motification-client: %v", err)
	}
	defer cleanup()
	log.Printf("Created motification-client")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var patientUid string
	{
		resp, err := client.RegisterPatient(ctx, &regprotobuf2.RegisterPatientRequest{
			Patient: &regprotobuf2.Patient{
				BSN:      bsn,
				FullName: name,
				Address: &regprotobuf2.Address{
					PostalCode:  "3731TB",
					HouseNumber: 79,
				},
				Contact: &regprotobuf2.Contact{
					EmailAddress: "mgrol@xebias.com",
					PhoneNumber:  "+31648928856",
				},
			},
		})
		if err != nil {
			log.Fatalf("sending sms: %s", err)
		}
		log.Printf("Patient registered: %+v", resp)
		patientUid = resp.PatientUid
	}
	{
		resp, err := client.CompletePatientRegistration(ctx, &regprotobuf2.CompletePatientRegistrationRequest{
			PatientUid: patientUid,
			Credentials: &regprotobuf2.RegistrationCredentials{
				Pincode: 1234,
			},
		})
		if err != nil {
			log.Fatalf("Completing patient registration failed: %s", err)
		}
		log.Printf("Patient completely registered: %+v", resp)
	}

}
