package notificationapi

import (
	"fmt"

	"google.golang.org/grpc"
)

const (
	DefaultPort = ":60003"
)

func NewGrpcClient(addressPort string) (NotificationClient, func(), error) {
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
	return NewNotificationClient(conn), cleanup, nil
}
