package appointmentapi

import (
	"fmt"

	"google.golang.org/grpc"
)

func NewExternalGrpcClient(addressPort string) (AppointmentExternalClient, func(), error) {
	// Prepare connection to the server.
	conn, err := grpc.Dial(addressPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, func() {}, fmt.Errorf("Error external-appointment-api-grpc-client: %v", err)
	}
	cleanup := func() {
		if conn != nil {
			conn.Close()
		}
	}
	return NewAppointmentExternalClient(conn), cleanup, nil
}

func NewInternalGrpcClient(addressPort string) (AppointmentInternalClient, func(), error) {
	// Prepare connection to the server.
	conn, err := grpc.Dial(addressPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, func() {}, fmt.Errorf("Error internal-appointment-api-grpc-client: %v", err)
	}
	cleanup := func() {
		if conn != nil {
			conn.Close()
		}
	}
	return NewAppointmentInternalClient(conn), cleanup, nil
}
