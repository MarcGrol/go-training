package main

import (
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/impl/datastoring"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/impl/emailsending"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/impl/pincoding"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/impl/smssending"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/impl/uuiding"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/regprotobuf"
	"log"
)

func main() {
	uuidGenerator := uuiding.New()
	patientStore := datastoring.New()
	pincoder := pincoding.New()
	emailSender := emailsending.New()
	smsSender := smssending.New()

	service := NewRegistrationService(uuidGenerator, patientStore, pincoder,
		emailSender, smsSender)
	err := StartGrpcServer(regprotobuf.DefaultPort, service)
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
