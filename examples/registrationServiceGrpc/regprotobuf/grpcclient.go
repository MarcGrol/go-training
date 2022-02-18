package regprotobuf

import (
	"fmt"
	"google.golang.org/grpc"
)

func NewGrpcClient(addressPort string) (RegistrationServiceClient, func(), error) {
	// Prepare connection to the server.
	conn, err := grpc.Dial(addressPort, grpc.WithInsecure()) //, grpc.WithBlock())
	if err != nil {
		return nil, func() {}, fmt.Errorf("Error creating registration grpc-client: %v", err)
	}
	cleanup := func() {
		if conn != nil {
			conn.Close()
		}
	}
	return NewRegistrationServiceClient(conn), cleanup, nil
}
