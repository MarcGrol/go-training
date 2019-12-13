package flightinfoapi

import (
	"fmt"

	"google.golang.org/grpc"
)

const (
	DefaultAddressPort = "localhost:16123"
)

func NewSyncGrpcClient(addressPort string) (FlightInfoClient, func(), error) {
	// Prepare connection to the server.
	conn, err := grpc.Dial(addressPort, grpc.WithInsecure()) //, grpc.WithBlock())
	if err != nil {
		return nil, func() {}, fmt.Errorf("Error starting flightinfoapi-sync-grpc-client: %s", err)
	}
	cleanup := func() {
		if conn != nil {
			conn.Close()
		}
	}
	return NewFlightInfoClient(conn), cleanup, nil
}

func NewAsyncGrpcClient(addressPort string) (FlightInfoAsyncClient, func(), error) {
	// Prepare connection to the server.
	conn, err := grpc.Dial(addressPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, func() {}, fmt.Errorf("Error flightinfoapi-async-grpc-client: %v", err)
	}
	cleanup := func() {
		if conn != nil {
			conn.Close()
		}
	}
	return NewFlightInfoAsyncClient(conn), cleanup, nil
}
