package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"

	pb "github.com/MarcGrol/go-training/examples/grpc/streaming/flightinfoapi"
	"google.golang.org/grpc"
)

type server struct {
	listener net.Listener
	pb.UnimplementedFlightInfoServer
	pb.UnimplementedFlightInfoAsyncServer
	flightStore  *flightStore
	sessionStore *sessionStore
}

func NewServer() *server {
	return &server{
		flightStore:  newFlightStore(),
		sessionStore: newSessionStore(),
	}
}

func (s *server) GRPCListenBlocking(port string) error {
	var err error
	log.Printf("Start listening on port %s", port)
	s.listener, err = net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen at port %s: %v", port, err)
	}
	log.Printf("Started listening on port %s", port)

	grpcServer := grpc.NewServer()
	pb.RegisterFlightInfoServer(grpcServer, s)
	pb.RegisterFlightInfoAsyncServer(grpcServer, s)

	// Start background routine that produces new flights that are passed to all connected client
	go simulateProductionOfFlightsInBackground(s.flightStore, s.sessionStore)

	log.Printf("Start serving flightinfo-server")
	err = grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *server) GetInfoOnFlight(ctx context.Context, req *pb.GetInfoOnFlightRequest) (*pb.GetFlightReply, error) {
	fmt.Printf("\nGetInfoOnFlight: %+v\n", req)

	return &pb.GetFlightReply{
		Flight: &pb.Flight{
			FlightUid:     uuid.New().String(),
			FlightNumber:  "KL1234",
			Direction:     pb.Direction_DEPARTURE,
			Date:          &pb.Date{Year: 2019, Month: 2, Day: 27},
			ScheduledTime: &pb.Time{Hour: 14, Minute: 14, Second: 47},
			Origin:        "AMS",
			Destination:   "LAX",
			Status:        pb.FlightStatus_DEPARTING,
		},
	}, nil
}

func (s *server) GetHistory(req *pb.HistoryRequest, stream pb.FlightInfoAsync_GetHistoryServer) error {
	for _, flight := range s.flightStore.getFlights() {
		// TODO do not ignore filter criteria in request
		if err := stream.Send(&flight); err != nil {
			return err
		}
		fmt.Printf("\nReturned historic flight: %s\n", flight.FlightNumber)
	}
	return nil // closes grpc session after return
}

func (s *server) SubscribeToEvents(stream pb.FlightInfoAsync_SubscribeToEventsServer) error {
	session, cleanup := newSession(stream)
	defer cleanup()

	s.sessionStore.register(session)
	defer s.sessionStore.unregister(session)

	return session.process() // closes grpc session after return
}
