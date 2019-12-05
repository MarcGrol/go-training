package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"
	"google.golang.org/grpc"
)

type server struct {
	listener net.Listener
	pb.UnimplementedNotificationServer
}

func newServer() *server {
	return &server{}
}

func (s *server) GRPCListenBlocking(port string) error {
	var err error
	s.listener, err = net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen at port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServer(grpcServer, newServer())

	log.Println("GRPPC server starts listening...")
	err = grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *server) SendEmail(ctx context.Context, in *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	if in == nil || in.GetEmail() == nil || in.GetEmail().GetRecipientEmailAddress() == "" || in.GetEmail().GetSubject() == "" {
		return &pb.SendEmailReply{
			Status: pb.DeliveryStatus_FAILED,
			Error: &pb.Error{
				Code:    400,
				Message: "Missing recipient-email-address or subject",
			},
		}, nil
	}

	// TODO call sendgrid or like
	log.Printf("Sent email to '%s' with subject '%s'", in.GetEmail().GetRecipientEmailAddress(), in.GetEmail().GetSubject())

	// Simulate error
	if strings.Contains(in.GetEmail().GetSubject(), "5") {
		return &pb.SendEmailReply{
			Status: pb.DeliveryStatus_FAILED,
			Error: &pb.Error{
				Code:    500,
				Message: "Internal error sending email",
			},
		}, nil
	}

	return &pb.SendEmailReply{Status: pb.DeliveryStatus_DELIVERED}, nil
}

func (s *server) SendSms(ctx context.Context, in *pb.SendSmsRequest) (*pb.SendSmsReply, error) {
	if in == nil || in.GetSms() == nil || in.GetSms().GetRecipientPhoneNumber() == "" || in.GetSms().GetBody() == "" {
		return &pb.SendSmsReply{
			Status: pb.DeliveryStatus_FAILED,
			Error: &pb.Error{
				Code:    400,
				Message: "Missing recipient-phone-number or body",
			},
		}, nil
	}

	// TODO call twilio or like
	log.Printf("Sent sms to '%s' with subject '%s'", in.GetSms().GetRecipientPhoneNumber(), in.GetSms().GetBody())

	return &pb.SendSmsReply{Status: pb.DeliveryStatus_DELIVERED}, nil
}
