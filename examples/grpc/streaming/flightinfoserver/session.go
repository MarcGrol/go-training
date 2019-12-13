package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"

	pb "github.com/MarcGrol/go-training/examples/grpc/streaming/flightinfoapi"
)

const (
	maxOutstandingCount = 100
)

type session struct {
	uid                 string
	internalChannel     chan pb.Flight
	externalChannel     chan pb.Acknowledgement
	quitChannel         chan bool
	stream              pb.FlightInfoAsync_KeepSynchronizingServer
	outstanding         map[string]bool
	maxOutstandingCount int
	sync.Mutex
}

func newSession(stream pb.FlightInfoAsync_KeepSynchronizingServer) (*session, func()) {
	sess := &session{
		uid:                 uuid.New().String(),
		internalChannel:     make(chan pb.Flight, maxOutstandingCount),
		externalChannel:     make(chan pb.Acknowledgement, maxOutstandingCount),
		quitChannel:         make(chan bool),
		stream:              stream,
		outstanding:         map[string]bool{},
		maxOutstandingCount: maxOutstandingCount,
	}
	cleanup := func() {
		log.Printf("Cleanup session: %s", sess.uid)
		close(sess.quitChannel)
		close(sess.internalChannel)
		close(sess.externalChannel)
	}
	log.Printf("Created session: %s", sess.uid)
	return sess, cleanup
}

func (s *session) getOutstandingCount() int {
	s.Lock()
	defer s.Unlock()

	return len(s.outstanding)
}

func (s *session) addOutstanding(uid string) {
	s.Lock()
	defer s.Unlock()

	s.outstanding[uid] = true
}

func (s *session) removeOutstanding(uid string) {
	s.Lock()
	defer s.Unlock()

	_, ok := s.outstanding[uid]
	if !ok {
		log.Printf("Remove non-existing event: %s", uid)
		return
	}
	delete(s.outstanding, uid)
}

func (s *session) process() error {
	go s.readExternal()

	return s.waitforEvents()
}

func (s *session) waitforEvents() error {
	log.Printf("Start processing events on session %s", s.uid)
	defer log.Printf("Done processing events on session %s", s.uid)

	heartbeatChannel := time.Tick(10 * time.Second)
	for {
		select {
		case <-heartbeatChannel:
			log.Printf("*** Send heartbeat on session %s\n", s.uid)
			err := s.stream.Send(&pb.FlightPdu{
				Payload: &pb.FlightPdu_Heartbeat{
					Heartbeat: &pb.Heartbeat{},
				},
			})
			if err != nil {
				log.Printf("*** Error writing heartbeat event: %s\n", err)
				return err
			}
		case <-s.quitChannel:
			log.Printf("*** Got quit signal\n")
			return nil
		case f := <-s.internalChannel:
			s.addOutstanding(f.FlightUid)
			err := s.stream.Send(&pb.FlightPdu{
				Payload: &pb.FlightPdu_Flight{
					Flight: &f,
				},
			})
			if err != nil {
				log.Printf("*** Error writing outbound event: %s\n", err)
				return err
			}
			fmt.Printf("\nWrote flight-event to remote client: %+v\n", f.FlightUid)
		case fc := <-s.externalChannel:
			s.removeOutstanding(fc.FlightUid)
			fmt.Printf("\nReceived confirmation from remote client: %+v\n", fc.FlightUid)
		}
	}
	return nil
}

func (s *session) readExternal() error {
	defer func() {
		s.quitChannel <- true
	}()

	log.Printf("Start reading external for session %s", s.uid)
	defer log.Printf("Done reading external for session %s", s.uid)

	for {
		fc, err := s.stream.Recv()
		if err != nil {
			log.Printf("*** Error reading from external: %s\n", err)
			return err
		}
		s.externalChannel <- *fc
	}
	return nil
}
