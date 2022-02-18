package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/regprotobuf"
	"log"
	"os"
	"time"
)

func main() {
	cliArgs := parseArgs()

	log.Printf("args: %+v", cliArgs)

	client, cleanup, err := regprotobuf.NewGrpcClient(regprotobuf.DefaultPort)
	if err != nil {
		log.Fatalf("*** Error creating motification-client: %v", err)
	}
	defer cleanup()
	log.Printf("Created registration-client")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if cliArgs.command == "start-registration" {
		patientUid, err := startRegistration(ctx, client, cliArgs.bsn, cliArgs.name, cliArgs.email)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Patient %s registered", patientUid)
	} else if cliArgs.command == "complete-registration" {
		err = completeRegistration(ctx, client, cliArgs.patientUid, cliArgs.pinCode)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Registration completed for patient %s", cliArgs.patientUid)
	} else if cliArgs.command == "bruteforce-registration" {
		err = bruteForceRegistration(ctx, client, cliArgs.patientUid)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Bruteforce completed for patient %s", cliArgs.patientUid)
	} else {
		log.Fatalf("Unrecognized command %s", cliArgs.command)
	}
}

func startRegistration(ctx context.Context, client regprotobuf.RegistrationServiceClient, bsn int, name, email string) (string, error) {
	resp, err := client.RegisterPatient(ctx, &regprotobuf.RegisterPatientRequest{
		Patient: &regprotobuf.Patient{
			BSN:      fmt.Sprintf("%d", bsn),
			FullName: name,
			Address: &regprotobuf.Address{
				PostalCode:  "3731TB",
				HouseNumber: 79,
			},
			Contact: &regprotobuf.Contact{
				EmailAddress: email,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("Error registering client: %s", err)
	}
	return resp.PatientUid, nil
}

func completeRegistration(ctx context.Context, client regprotobuf.RegistrationServiceClient, patientUid string, pinCode int) error {
	_, err := client.CompletePatientRegistration(ctx, &regprotobuf.CompletePatientRegistrationRequest{
		PatientUid: patientUid,
		Credentials: &regprotobuf.RegistrationCredentials{
			Pincode: int32(pinCode),
		},
	})
	if err != nil {
		return fmt.Errorf("Error completing patient registration: %s", err)
	}

	return nil
}

func bruteForceRegistration(ctx context.Context, client regprotobuf.RegistrationServiceClient, patientUid string) error {
	for i := 1; i <= 10; i++ {
		log.Printf("Start guessing pincode %d", i)
		_, err := client.CompletePatientRegistration(ctx, &regprotobuf.CompletePatientRegistrationRequest{
			PatientUid: patientUid,
			Credentials: &regprotobuf.RegistrationCredentials{
				Pincode: int32(i),
			},
		})
		if err != nil {
			log.Printf("Error completing patient registration: %s", err)
		} else {
			log.Printf("Pincode successfully guessed")
			return nil
		}
	}
	return fmt.Errorf("Error guessing pin-code")
}

type args struct {
	command    string
	email      string
	name       string
	bsn        int
	pinCode    int
	patientUid string
}

func parseArgs() args {

	help := flag.Bool("help", false, "This help text")
	command := flag.String("command", "start-registration", "Command (start-registration, complete-registration or brute-force-registration)")

	bsn := flag.Int("bsn", 12345, "Bsn number of patient")
	name := flag.String("name", "Michael Jordan", "Name of patient")
	email := flag.String("email", "eva.marc@hetnet.nl", "Email address of patient")

	patientUid := flag.String("patient-uid", "", "Uid of patient")
	pinCode := flag.Int("pin-code", -1, "Pincode to confirm registration")

	flag.Parse()

	if help != nil && *help {
		fmt.Fprintf(os.Stderr, "\nUsage:\n")
		fmt.Fprintf(os.Stderr, "\t[flags]\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(-1)
	}

	return args{
		command: *command,
		email:   *email,
		name:    *name,
		bsn:     *bsn,

		pinCode:    *pinCode,
		patientUid: *patientUid,
	}
}
