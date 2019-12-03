package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

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

func (s *server)GRPCListenBlocking(port string) error {
	var err error
	s.listener, err = net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen at port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServer(grpcServer, New())

	log.Println("GRPPC server starts listening...")
	err = grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *server)ListenHttpBlocking(grpcPort string, restPort string) error{
	// start GRPC server in the background
	go func() {
		err := s.GRPCListenBlocking(grpcPort)
		if err != nil {
			log.Fatalf("Error starting grpc-notification server: %s", err)
		}
	}()

	// give it some time to startup
	time.Sleep(2* time.Second)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the just created GRPC server
	conn, err := grpc.Dial(grpcPort, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("fail to dial: %v", err)
	}
	defer conn.Close()

	// Register REST gateway that proxies rest to grpc and back
	rmux := runtime.NewServeMux()
	client := pb.NewNotificationClient(conn)
	err = pb.RegisterNotificationHandlerClient(ctx, rmux, client)
	if err != nil {
		return fmt.Errorf("Error registering http-client: %s", err)
	}

	// Serve the swagger-ui and swagger file
	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	mux.HandleFunc("/swagger.json", serveSwagger)
	fs := http.FileServer(http.Dir("swaggerui"))
	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui", fs))

	log.Println("REST server starts listening...")
	err = http.ListenAndServe(restPort, mux)
	if err != nil {
		return fmt.Errorf("Error listening on http: %s", err)
	}
	return nil
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swaggerui/swagger.json")
}

func (s *server) SendEmail(ctx context.Context, in *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	// TODO check mandatory parameters
	log.Printf("Send email to '%s' with subject '%s'", in.GetEmail().GetRecipientEmailAddress(), in.GetEmail().GetSubject())
	if strings.Contains(in.GetEmail().GetSubject(), "5") {
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
	// TODO check mandatory parameters
	log.Printf("Send sms to '%s' with body '%s'", in.GetSms().GetRecipientPhoneNumber(), in.GetSms().GetBody())
	msgUID := uuid.New().String()
	status := pb.NotificationStatus_PENDING
	s.messages[msgUID] = status
	return &pb.SendSmsReply{NotificationUid:msgUID}, nil
}

func (s *server) GetNotificationStatus(ctx context.Context, in *pb.GetNotificationStatusRequest) (*pb.GetNotificationStatusReply, error) {
	// TODO check mandatory parameters
	log.Printf("Get status of notification with uid '%s'", in.GetNotificationUid())
	status, found := s.messages[in.GetNotificationUid()]
	if !found {
		return &pb.GetNotificationStatusReply{
			Status:pb.NotificationStatus_UNKNOWN,
			Error: &pb.Error{
				Code:                 404,
				Message:              "Notification with uid not found",
			},
		}, nil
	}
	return &pb.GetNotificationStatusReply{
		Status: status,
	}, nil
}
