package main

import (
	"context"
	"fmt"
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

	log.Println("GRPPC server starts listening...")
	err = grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *server) GetPatientOnUid(ctx context.Context, in *pb.GetPatientOnUidRequest) (*pb.GetPatientOnUidReply, error) {
	// Validate input
	if in == nil || in.GetPatientUid() == "" {
		return &pb.GetPatientOnUidReply{
			Error: &pb.Error{
				Code:    400,
				Message: "Missing patient-uid",
			},
		}, nil
	}

	// Perform lookup
	patient, found, err := s.patientStore.GetPatientOnUid(in.PatientUid)
	if err != nil {
		return &pb.GetPatientOnUidReply{
			Error: &pb.Error{
				Code:    500,
				Message: "Technical error",
				Details: err.Error(),
			},
		}, nil
	}
	if !found {
		return &pb.GetPatientOnUidReply{
			Error: &pb.Error{
				Code:    404,
				Message: "Patient with uid not found",
			},
		}, nil
	}

	log.Printf("Patient with uid %s found: %_v", in.PatientUid, patient)

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
