package main

import (
	"fmt"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/regprotobuf"

	"google.golang.org/grpc"
)

func NewGrpcClient(addressPort string) (regprotobuf.RegistrationServiceClient, func(), error) {
	// Prepare connection to the server.
	conn, err := grpc.Dial(addressPort, grpc.WithInsecure()) //, grpc.WithBlock())
	if err != nil {
		return nil, func() {}, fmt.Errorf("Error creating notification-api-grpc-client: %v", err)
	}
	cleanup := func() {
		if conn != nil {
			conn.Close()
		}
	}
	return regprotobuf.NewRegistrationServiceClient(conn), cleanup, nil
}
