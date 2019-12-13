package main

import (
	"fmt"
	"log"
	"time"

	pb "github.com/MarcGrol/go-training/examples/grpc/streaming/flightinfoapi"
	"github.com/google/uuid"
)

func simulateProductionOfFlightsInBackground(flightStore *flightStore, sessionStore *sessionStore) {
	log.Printf("Start producing")
	idx := 0
	tick := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-tick:
			for _, sess := range sessionStore.getSessions() {
				outstandingCount := sess.getOutstandingCount()
				if outstandingCount >= sess.maxOutstandingCount {
					// prevent that slow client also delays other clients
					log.Printf("Session %s exceeds max-outstanding count", sess.uid)
				} else {
					now := time.Now()
					f := pb.Flight{
						FlightUid:     uuid.New().String(),
						FlightNumber:  fmt.Sprintf("KL%05d", idx),
						Direction:     pb.Direction_ARRIVAL,
						Date:          &pb.Date{Year: int32(now.Year()), Month: int32(now.Month()), Day: int32(now.Day())},
						ScheduledTime: &pb.Time{Hour: int32(now.Hour()), Minute: int32(now.Minute()), Second: int32(now.Second())},
						Origin:        "AMS",
						Destination:   "LAX",
						Status:        pb.FlightStatus_LANDING,
					}

					sess.internalChannel <- f

					// allow late arrivers to read history
					flightStore.addFlight(f)
				}
			}
			idx++
		}
	}
}
