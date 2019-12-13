package main

import (
	"log"
	"time"

	pb "github.com/MarcGrol/go-training/examples/grpc/streaming/flightinfoapi"
)

const (
	port = ":50051"
)

func main() {
	s := New()
	time.Sleep(time.Second * 2)
	err := s.GRPCListenBlocking(pb.DefaultAddressPort)
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
