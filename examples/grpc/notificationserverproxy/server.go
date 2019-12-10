package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "github.com/MarcGrol/go-training/examples/grpc/notificationapi"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type server struct {
}

func New() *server {
	return &server{}
}

func (s *server) ListenHttpBlocking(grpcPort string, restPort string) error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	// create a regular grpc-client
	grpcClient, cleanup, err := pb.NewGrpcClient(grpcPort)
	if err != nil {
		return err
	}
	defer cleanup()

	// create a rest-endpoint that uses the regular grpc-client to forward to the real grpc server
	rmux := runtime.NewServeMux()
	err = pb.RegisterNotificationHandlerClient(ctx, rmux, grpcClient)
	if err != nil {
		return fmt.Errorf("Error registering notif-http-client: %s", err)
	}
	mux.Handle("/", rmux)

	// Serve the swagger-ui and swagger file
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "swaggerui/swagger.json")
	})
	fs := http.FileServer(http.Dir("swaggerui"))
	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))

	log.Printf("REST server starts listening on port %s\n", restPort)
	err = http.ListenAndServe(restPort, mux)
	if err != nil {
		return fmt.Errorf("Error listening on http: %s", err)
	}

	return nil
}
