package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"
)

type server struct {
	listener net.Listener
	pb.UnimplementedPatientInfoServer
	patientStore PatientStore
}

func newServer(store PatientStore) *server {
	return &server{
		patientStore: store,
	}
}

func (s *server) GRPCListenBlocking(port string) error {
	var err error
	s.listener, err = net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen at port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPatientInfoServer(grpcServer, s)

	log.Printf("GRPPC server starts listening at port %s...", port)
	err = grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *server) GetPatientOnUid(ctx context.Context, in *pb.GetPatientOnUidRequest) (*pb.GetPatientOnUidReply, error) {
	// Validate input
	if in == nil || in.GetPatientUid() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing patient-uid")
	}

	// Perform lookup
	patient, found, err := s.patientStore.GetPatientOnUid(in.PatientUid)
	if err != nil {
		status.Errorf(codes.Internal, "Technical error: %s", err)
	}
	if !found {
		status.Errorf(codes.NotFound, "Patient with uid not found")
	}

	log.Printf("Patient with uid %s found: %+v", in.PatientUid, patient)

	// return response
	return &pb.GetPatientOnUidReply{
		Patient: &pb.Patient{
			Uid:          patient.UID,
			FullName:     patient.FullName,
			AddressLine:  patient.AddressLine,
			PhoneNumber:  patient.PhoneNumber,
			EmailAddress: patient.EmailAddress,
		},
	}, nil
}
