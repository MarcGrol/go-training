package main

import (
	"sync"

	pb "github.com/MarcGrol/go-training/examples/grpc/streaming/flightinfoapi"
)

type flightStore struct {
	flights []pb.Flight
	sync.Mutex
}

func newFlightStore() *flightStore {
	return &flightStore{
		flights: []pb.Flight{},
	}
}

func (fs *flightStore) addFlight(f pb.Flight) {
	fs.Lock()
	defer fs.Unlock()

	fs.flights = append(fs.flights, f)
}

func (fs *flightStore) getFlight(flightNumber string, direction pb.Direction, date pb.Date) []pb.Flight {
	fs.Lock()
	defer fs.Unlock()

	return fs.flights
}

func (fs flightStore) getFlights() []pb.Flight {
	fs.Lock()
	defer fs.Unlock()

	return fs.flights
}
