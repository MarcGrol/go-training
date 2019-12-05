package patientinfoapi

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
)

const (
	DefaultPort = ":60002"
)

func NewGrpcClient(addressPort string) (PatientInfoClient, func(), error) {
	// Prepare connection to the server.
	conn, err := grpc.Dial(addressPort, grpc.WithInsecure()) //), grpc.WithBlock())
	if err != nil {
		log.Printf("Error creating patient-client:%s", err)
		return nil, func() {}, fmt.Errorf("Error creating patientinfo-grpc-client: %v", err)
	}
	cleanup := func() {
		if conn != nil {
			conn.Close()
		}
	}
	return NewPatientInfoClient(conn), cleanup, nil
}
