package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/MarcGrol/go-training/examples/grpc/streaming/flightinfoapi"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 365*24*time.Hour)
	defer cancel()

	synchronousInteraction(ctx)

	asyncStreamingInteraction(ctx)
}

func synchronousInteraction(ctx context.Context) error {
	client, cleanup, err := pb.NewSyncGrpcClient(pb.DefaultAddressPort)
	if err != nil {
		log.Printf("*** Error ccreating client: %v", err)
		return err
	}
	defer cleanup()

	resp, err := client.GetInfoOnFlight(ctx, &pb.GetInfoOnFlightRequest{
		Date:         &pb.Date{Year: 2019, Month: 2, Day: 27},
		Direction:    pb.Direction_DEPARTURE,
		FlightNumber: "KL1234",
	})
	if err != nil {
		log.Printf("Error getting flight-info:%s", err)
		return err
	}
	if resp.Error != nil {
		log.Printf("Error getting flight-info:%+v", resp.Error)
	} else {
		fmt.Printf("\nsync-resp:%+s\n", resp.Flight.FlightNumber)
	}
	return nil
}

func asyncStreamingInteraction(ctx context.Context) error {
	client, cleanup, err := pb.NewAsyncGrpcClient(pb.DefaultAddressPort)
	if err != nil {
		log.Printf("*** Error creating client: %v\n", err)
		return err
	}
	defer cleanup()

	err = oneWayStreaming(ctx, client)
	if err != nil {
		log.Printf("*** Error performing one-way streaming: %v\n", err)
		return err
	}

	err = twoWayStreaming(ctx, client)
	if err != nil {
		log.Printf("*** Error performing two-way streaming: %v\n", err)
		return err
	}
	return nil
}

func oneWayStreaming(ctx context.Context, client pb.FlightInfoAsyncClient) error {
	stream, err := client.GetHistory(ctx, &pb.HistoryRequest{
		StartDate: &pb.Date{Year: 2019, Month: 2, Day: 27},
	})
	if err != nil {
		log.Printf("Error getting flight-info history:%s\n", err)
		return err
	}

	for {
		flight, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Error getting history-request:%s\n", err)
			return err
		}
		fmt.Printf("\nhistoric flight:%+s\n", flight.FlightNumber)
	}
	log.Printf("Done fetching history\n\n")
	return nil
}

func twoWayStreaming(ctx context.Context, client pb.FlightInfoAsyncClient) error {
	stream, err := client.KeepSynchronizing(ctx)
	if err != nil {
		log.Printf("Error getting flight-info:%s\n", err)
		return err
	}

	log.Printf("Waitfor events on\n")
	for {
		pdu, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("No more events: %s\n", err)
				break
			}
			log.Printf("Error subscribing to flight-info events:%s\n", err)
			return err
		}

		if pdu.GetHeartbeat() != nil {
			log.Printf("Got heartbeat\n")
		} else if pdu.GetFlight() != nil {

			fmt.Printf("\nGet recent flight info:%+s\n", pdu.GetFlight().GetFlightNumber())

			// ackknowledge receipt back to server
			err = stream.Send(&pb.Acknowledgement{
				FlightUid: pdu.GetFlight().FlightUid,
			})
			if err != nil {
				if err == io.EOF {
					log.Printf("No more events: %s\n", err)
					break
				}
				log.Printf("Error confirming flight-info event:%s\n", pdu.GetFlight().FlightUid)
				return err
			}
		}
	}
	return nil
}
