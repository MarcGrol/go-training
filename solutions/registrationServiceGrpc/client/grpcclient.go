package main

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/regprotobuf"
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
