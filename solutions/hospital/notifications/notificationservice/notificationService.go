package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"
)

type EmailSender interface {
	Send(recipientEmail, subject, emailBody string) error
}

type SmsSender interface {
	Send(recipientPhoneNumber, messageBody string) error
}

type server struct {
	listener net.Listener
	pb.UnimplementedNotificationServer

	emailSender EmailSender
	smsSender   SmsSender
}

func newServer(emailSender EmailSender, smsSender SmsSender) *server {
	return &server{
		emailSender: emailSender,
		smsSender:   smsSender,
	}
}

func (s *server) GRPCListenBlocking(port string) error {
	var err error
	s.listener, err = net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen at port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServer(grpcServer, s)

	log.Printf("GRPPC server starts listening at port %s...", port)
	err = grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *server) SendEmail(ctx context.Context, in *pb.SendEmailRequest) (*pb.SendReply, error) {
	if in == nil || in.GetEmail() == nil || in.GetEmail().GetRecipientEmailAddress() == "" || in.GetEmail().GetSubject() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing recipient-email-address or subject")
	}

	err := s.emailSender.Send(in.GetEmail().GetRecipientEmailAddress(), in.GetEmail().GetSubject(), in.GetEmail().GetBody())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error sending out email:%s", err)
	}
	log.Printf("Sent email to '%s' with subject '%s'", in.GetEmail().GetRecipientEmailAddress(), in.GetEmail().GetSubject())

	return &pb.SendReply{Status: pb.DeliveryStatus_DELIVERED}, nil
}

func (s *server) SendSms(ctx context.Context, in *pb.SendSmsRequest) (*pb.SendReply, error) {
	if in == nil || in.GetSms() == nil || in.GetSms().GetRecipientPhoneNumber() == "" || in.GetSms().GetBody() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing recipient-phone-number or body")
	}

	err := s.smsSender.Send(in.GetSms().GetRecipientPhoneNumber(), in.GetSms().GetBody())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error sending out sms:%s", err)
	}
	log.Printf("Sent sms to '%s' with body '%s'", in.GetSms().GetRecipientPhoneNumber(), in.GetSms().GetBody())

	return &pb.SendReply{Status: pb.DeliveryStatus_DELIVERED}, nil
}
