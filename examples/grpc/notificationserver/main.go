
package main

import (
	_ "github.com/MarcGrol/go-training/examples/grpc/notifapi"
	"log"
	"time"
)

func main() {
	s := New()
	time.Sleep(time.Second*2)
	err := s.GRPCListenBlocking(":50051")
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}

