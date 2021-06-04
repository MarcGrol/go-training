package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	pb "github.com/MarcGrol/go-training/examples/grpc/notification/notificationapi"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	listener net.Listener
	pb.UnimplementedNotificationServer
	messages map[string]pb.NotificationStatus
}

func New(port string) (*service, error) {
	notifServer := &service{
		messages: map[string]pb.NotificationStatus{},
	}

	var err error
	notifServer.listener, err = net.Listen("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("failed to listen at port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServer(grpcServer, notifServer)

	log.Printf("GRPPC service starts listening on port %s", port)
	err = grpcServer.Serve(notifServer.listener)
	if err != nil {
		return nil, fmt.Errorf("failed to serve: %v", err)
	}
	return notifServer, nil
}

func (s *service) SendEmail(ctx context.Context, in *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	// TODO check mandatory parameters
	if strings.Contains(in.GetEmail().GetSubject(), "bad request") {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}
	if strings.Contains(in.GetEmail().GetSubject(), "not found") {
		return nil, status.Error(codes.NotFound, "Not found")
	}
	if strings.Contains(in.GetEmail().GetSubject(), "internal error") {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	if strings.Contains(in.GetEmail().GetSubject(), "permission denied") {
		return nil, status.Error(codes.PermissionDenied, "Permission denied")
	}
	log.Printf("Send email to '%s' with subject '%s'", in.GetEmail().GetRecipientEmailAddress(), in.GetEmail().GetSubject())
	msgUID := uuid.New().String()
	status := pb.NotificationStatus_SENT
	s.messages[msgUID] = status
	return &pb.SendEmailReply{NotificationUid: msgUID}, nil
}

func (s *service) SendSms(ctx context.Context, in *pb.SendSmsRequest) (*pb.SendSmsReply, error) {
	// TODO check mandatory parameters
	if strings.Contains(in.GetSms().GetBody(), "bad request") {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}
	if strings.Contains(in.GetSms().GetBody(), "not found") {
		return nil, status.Error(codes.NotFound, "Not found")
	}
	if strings.Contains(in.GetSms().GetBody(), "internal error") {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	if strings.Contains(in.GetSms().GetBody(), "permission denied") {
		return nil, status.Error(codes.PermissionDenied, "Permission denied")
	}
	log.Printf("Send sms to '%s' with body '%s'", in.GetSms().GetRecipientPhoneNumber(), in.GetSms().GetBody())
	msgUID := uuid.New().String()
	status := pb.NotificationStatus_PENDING
	s.messages[msgUID] = status
	return &pb.SendSmsReply{NotificationUid: msgUID}, nil
}

func (s *service) GetNotificationStatus(ctx context.Context, in *pb.GetNotificationStatusRequest) (*pb.GetNotificationStatusReply, error) {
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
