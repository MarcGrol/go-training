package main

import (
	"context"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/regprotobuf"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalf("Usage: %s bsn name email phone", os.Args[0])
	}
	bsn := os.Args[1]
	name := os.Args[2]
	email := os.Args[3]
	phone := os.Args[4]

	client, cleanup, err := regprotobuf.NewGrpcClient(regprotobuf.DefaultPort)
	if err != nil {
		log.Fatalf("*** Error creating motification-client: %v", err)
	}
	defer cleanup()
	log.Printf("Created registration-client")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var patientUid string
	{
		resp, err := client.RegisterPatient(ctx, &regprotobuf.RegisterPatientRequest{
			Patient: &regprotobuf.Patient{
				BSN:      bsn,
				FullName: name,
				Address: &regprotobuf.Address{
					PostalCode:  "3731TB",
					HouseNumber: 79,
				},
				Contact: &regprotobuf.Contact{
					EmailAddress: email,
					PhoneNumber:  phone,
				},
			},
		})
		if err != nil {
			log.Fatalf("Error registering client: %s", err)
		}
		log.Printf("Patient registered: %+v", resp)
		patientUid = resp.PatientUid
	}
	for i := 1; i <= 10; i++ {
		log.Printf("Start patient compltion with pin %d", i)
		resp, err := client.CompletePatientRegistration(ctx, &regprotobuf.CompletePatientRegistrationRequest{
			PatientUid: patientUid,
			Credentials: &regprotobuf.RegistrationCredentials{
				Pincode: int32(i),
			},
		})
		if err != nil {
			log.Printf("Error completing patient registration: %s", err)
		}
		log.Printf("Patient completely registered: %+v", resp)
	}

}
