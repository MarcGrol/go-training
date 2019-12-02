package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	pb "github.com/MarcGrol/go-training/examples/project/notificationService/spec"
	"github.com/google/uuid"

)

type server struct {
	listener net.Listener
	pb.UnimplementedNotificationServer
	messages map[string]pb.NotificationStatus
}

func New() *server{
	return &server{
		messages:map[string]pb.NotificationStatus{},
	}
}

func (s *server)ListenBlocking(port string) error {
	var err error
	s.listener, err = net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen at port %s: %v", port, err)
	}
	
	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServer(grpcServer, New())
	err = grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *server) SendEmail(ctx context.Context, in *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	log.Printf("Send email to %s with subject %s", in.GetRecipientEmailAddress(), in.GetSubject())
	if strings.Contains(in.GetSubject(), "5") {
		return &pb.SendEmailReply{
			NotificationUid:"",
			Error: &pb.Error{
				Code:    500,
				Message: "Internal error sending email",
			},
		}, nil
	}
	msgUID := uuid.New().String()
	status := pb.NotificationStatus_SENT
	s.messages[msgUID] = status
	return &pb.SendEmailReply{NotificationUid:msgUID}, nil
}

func (s *server) SendSms(ctx context.Context, in *pb.SendSmsRequest) (*pb.SendSmsReply, error) {
	log.Printf("Send sms to %s with subject %s", in.GetRecipientPhoneNumber(), in.GetBody())
	msgUID := uuid.New().String()
	status := pb.NotificationStatus_PENDING
	s.messages[msgUID] = status
	return &pb.SendSmsReply{NotificationUid:msgUID}, nil
}

func (s *server) GetNotificationStatus(ctx context.Context, in *pb.GetNotificationStatusRequest) (*pb.GetNotificationStatusReply, error) {
	log.Printf("Get status of notification with uid %s", in.GetNotificationUid())
	status := s.messages[in.GetNotificationUid()]
	return &pb.GetNotificationStatusReply{Status:status}, nil
}
