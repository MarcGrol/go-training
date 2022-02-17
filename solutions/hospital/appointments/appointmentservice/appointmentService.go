package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentservice/appointmentstore"

	"google.golang.org/grpc"

	pb "github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentapi"
	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"
	"github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"
)

type server struct {
	listener net.Listener
	pb.UnimplementedAppointmentExternalServer
	pb.UnimplementedAppointmentInternalServer

	appointmentStore   appointmentstore.AppointmentStore
	patientInfoClient  patientinfoapi.PatientInfoClient
	notificationClient notificationapi.NotificationClient
}

func newServer(store appointmentstore.AppointmentStore, patientInfoClient patientinfoapi.PatientInfoClient, notificationClient notificationapi.NotificationClient) *server {
	return &server{
		appointmentStore:   store,
		patientInfoClient:  patientInfoClient,
		notificationClient: notificationClient,
	}
}

func (s *server) GRPCListenBlocking(port string) error {
	var err error
	s.listener, err = net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen at port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAppointmentExternalServer(grpcServer, s)
	pb.RegisterAppointmentInternalServer(grpcServer, s)

	log.Printf("GRPPC server starts listening at port %s...", port)
	err = grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *server) GetAppointmentsOnUser(c context.Context, in *pb.GetAppointmentsOnUserRequest) (*pb.GetAppointmentsReply, error) {
	// Validate input
	if in == nil || in.GetUserUid() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing user-uid")
	}

	// Perform lookup
	internalAppointments, err := s.appointmentStore.GetAppointmentsOnUserUid(c, in.UserUid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Technical error fetching appointments on user:%s", err)
	}

	return returnAppointmentList(internalAppointments), nil
}

func (s *server) GetAppointmentsOnStatus(c context.Context, in *pb.GetAppointmentsOnStatusRequest) (*pb.GetAppointmentsReply, error) {
	// Validate input
	if in == nil || in.GetStatus() == pb.AppointmentStatus_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid status")
	}

	// Perform lookup
	internalAppointments, err := s.appointmentStore.GetAppointmentsOnStatus(c, appointmentstore.AppointmentStatus(in.Status))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Technical error fetching appointments on status:%s", err)
	}

	return returnAppointmentList(internalAppointments), nil
}

func (s *server) RequestAppointment(c context.Context, in *pb.RequestAppointmentRequest) (*pb.AppointmentReply, error) {
	// Validate input
	if in == nil || in.GetAppointment() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Missing appointment")
	}
	if in.GetAppointment().GetUserUid() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing userUid")
	}
	if in.GetAppointment().GetDateTime() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing dateTime")
	}
	if in.GetAppointment().GetLocation() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing location")
	}
	if in.GetAppointment().GetTopic() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing topic")
	}

	// Check if patient exists
	_, err := s.patientInfoClient.GetPatientOnUid(c, &patientinfoapi.GetPatientOnUidRequest{PatientUid: in.Appointment.UserUid})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Technical error fetching patient on uid:%s", err)
	}

	// Adjust datastore
	appointmentCreated, err := s.appointmentStore.PutAppointment(c, convertIntoInternal(*in.Appointment))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Technical error creating appointment:%s", err)
	}
	log.Printf("Persisted new appointment")

	return returnSingleAppointment(appointmentCreated), nil
}

func (s *server) ModifyAppointmentStatus(c context.Context, in *pb.ModifyAppointmentStatusRequest) (*pb.AppointmentReply, error) {
	// Validate input
	// Validate input
	if in == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Missing request")
	}

	if in.GetAppointmentUid() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing appointmentUid")
	}
	if in.GetStatus() == pb.AppointmentStatus_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "Missing status")
	}

	// Perform lookup
	internalAppointment, found, err := s.appointmentStore.GetAppointmentOnUid(c, in.AppointmentUid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error fetching appointment on uid:%s", err)
	}
	if !found {
		return nil, status.Errorf(codes.NotFound, "Appointment with uid not found")
	}
	log.Printf("Got appointment:%+v", internalAppointment)

	// Fetch patient details
	getPatientOnUidResp, err := s.patientInfoClient.GetPatientOnUid(c, &patientinfoapi.GetPatientOnUidRequest{PatientUid: internalAppointment.UserUID})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Technical error fetching patient on uid:%s", err)
	}

	log.Printf("Got patient:%+v", getPatientOnUidResp.Patient)

	{ // Send out email
		sendEmailResponse, err := s.notificationClient.SendEmail(c, &notificationapi.SendEmailRequest{
			Email: &notificationapi.EmailMessage{
				RecipientEmailAddress: getPatientOnUidResp.Patient.EmailAddress,
				Subject:               "Appointment confirmed",              // TODO use template
				Body:                  "Appointment confirmed with details", // TODO use template
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Technical error sending email:%s", err)
		}
		log.Printf("Send email:%+v", sendEmailResponse)
	}

	{
		// Send out sms
		sendSmsResponse, err := s.notificationClient.SendSms(c, &notificationapi.SendSmsRequest{
			Sms: &notificationapi.SmsMessage{
				RecipientPhoneNumber: getPatientOnUidResp.Patient.PhoneNumber,
				Body:                 "Appointment confirmed", // TODO use template
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Technical error sending sms:%s", err)
		}
		log.Printf("Send sms:%+v", sendSmsResponse)
	}

	// Adjust datastore
	internalAppointment.Status = appointmentstore.AppointmentStatusConfirmed
	appointmentAdjusted, err := s.appointmentStore.PutAppointment(c, internalAppointment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error persisting modified appointment:%s", err)
	}
	log.Printf("Persisted adjusted appointment: %+v", appointmentAdjusted)

	return returnSingleAppointment(appointmentAdjusted), nil
}

func returnAppointmentList(internalAppointments []appointmentstore.Appointment) *pb.GetAppointmentsReply {
	externalAppointments := []*pb.Appointment{}
	for _, a := range internalAppointments {
		externalAppointments = append(externalAppointments, convertIntoExternal(a))
	}
	return &pb.GetAppointmentsReply{
		Appointments: externalAppointments,
	}
}

func returnSingleAppointment(internalAppointmnent appointmentstore.Appointment) *pb.AppointmentReply {
	return &pb.AppointmentReply{
		Appointment: convertIntoExternal(internalAppointmnent),
	}
}

func convertIntoExternal(a appointmentstore.Appointment) *pb.Appointment {
	return &pb.Appointment{
		AppointmentUid: a.AppointmentUID,
		UserUid:        a.UserUID,
		DateTime:       a.DateTime,
		Location:       a.Location,
		Topic:          a.Topic,
		Status:         pb.AppointmentStatus(a.Status),
	}
}

func convertIntoInternal(a pb.Appointment) appointmentstore.Appointment {
	return appointmentstore.Appointment{
		AppointmentUID: a.AppointmentUid,
		UserUID:        a.UserUid,
		DateTime:       a.DateTime,
		Location:       a.Location,
		Topic:          a.Topic,
		Status:         appointmentstore.AppointmentStatus(a.Status),
	}
}
