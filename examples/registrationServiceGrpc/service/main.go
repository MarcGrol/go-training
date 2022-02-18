package main

import (
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/impl/datastoring"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/impl/emailsending"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/impl/pincoding"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/impl/smssending"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/impl/uuiding"
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
	err := regprotobuf.StartGrpcServer(regprotobuf.DefaultPort, service)
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
