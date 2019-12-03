package main

import (
	"context"
	"fmt"
	pb "github.com/MarcGrol/go-training/examples/grpc/spec"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type server struct {
}

func New() *server{
	return &server{}
}

func (s *server)ListenHttpBlocking(grpcPort string, restPort string) error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	{
		// Connect to the just created GRPC server
		conn, err := grpc.Dial(grpcPort, grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("fail to dial: %v", err)
		}
		defer conn.Close()

		// Register REST gateway that proxies rest to grpc and back
		client := pb.NewNotificationClient(conn)
		rmux := runtime.NewServeMux()
		err = pb.RegisterNotificationHandlerClient(ctx, rmux, client)
		if err != nil {
			return fmt.Errorf("Error registering http-client: %s", err)
		}
		mux.Handle("/", rmux)
	}

	{
		// Serve the swagger-ui and swagger file
		mux.HandleFunc("/swagger.json", serveSwagger)
		fs := http.FileServer(http.Dir("swaggerui"))
		mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui", fs))
	}

	log.Printf("REST server starts listening...\n")
	err := http.ListenAndServe(restPort, mux)
	if err != nil {
		return fmt.Errorf("Error listening on http: %s", err)
	}

	return nil
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swaggerui/swagger.json")
}
