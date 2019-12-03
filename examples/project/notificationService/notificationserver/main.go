
package main

import (
	"log"
	_ "github.com/MarcGrol/go-training/examples/project/notificationService/spec"
	"time"
)

func main() {
	s := New()
	time.Sleep(time.Second*2)
	err := s.ListenHttpBlocking(":50051", ":8080")
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}

