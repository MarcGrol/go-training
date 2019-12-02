
package main

import (
	"log"
	_ "github.com/MarcGrol/go-training/examples/project/notificationService/spec"
)

func main() {
	s := New()
	err := s.ListenBlocking(":50051")
	if err != nil {
		log.Fatalf("Error starting notification server: %s", err)
	}
}
