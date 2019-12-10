package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/MarcGrol/go-training/examples/grpc/notificationapi"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	listener net.Listener
	pb.UnimplementedNotificationServer
	messages map[string]pb.NotificationStatus
}

func New() *server {
	return &server{
		messages: map[string]pb.NotificationStatus{},
	}
}

func (s *server) GRPCListenBlocking(port string) error {
	var err error
	s.listener, err = net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen at port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServer(grpcServer, New())

	log.Printf("GRPPC server starts listening on port %s", port)
	err = grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *server) SendEmail(ctx context.Context, in *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	// TODO check mandatory parameters
	log.Printf("Send email to '%s' with subject '%s'", in.GetEmail().GetRecipientEmailAddress(), in.GetEmail().GetSubject())
	if strings.Contains(in.GetEmail().GetSubject(), "5") {
		return &pb.SendEmailReply{
			NotificationUid: "",
			Error: &pb.Error{
				Code:    500,
				Message: "Internal error sending email",
			},
		}, nil
	}
	msgUID := uuid.New().String()
	status := pb.NotificationStatus_SENT
	s.messages[msgUID] = status
	return &pb.SendEmailReply{NotificationUid: msgUID}, nil
}

func (s *server) SendSms(ctx context.Context, in *pb.SendSmsRequest) (*pb.SendSmsReply, error) {
	// TODO check mandatory parameters
	log.Printf("Send sms to '%s' with body '%s'", in.GetSms().GetRecipientPhoneNumber(), in.GetSms().GetBody())
	msgUID := uuid.New().String()
	status := pb.NotificationStatus_PENDING
	s.messages[msgUID] = status
	return &pb.SendSmsReply{NotificationUid: msgUID}, nil
}

func (s *server) GetNotificationStatus(ctx context.Context, in *pb.GetNotificationStatusRequest) (*pb.GetNotificationStatusReply, error) {
	// TODO check mandatory parameters
	log.Printf("Get status of notification with uid '%s'", in.GetNotificationUid())
	status, found := s.messages[in.GetNotificationUid()]
	if !found {
		return &pb.GetNotificationStatusReply{
			Status: pb.NotificationStatus_UNKNOWN,
			Error: &pb.Error{
				Code:    404,
				Message: "Notification with uid not found",
			},
		}, nil
	}
	return &pb.GetNotificationStatusReply{
		Status: status,
	}, nil
}
