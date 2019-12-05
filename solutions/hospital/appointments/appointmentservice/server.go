package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentapi"

	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"
	"github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"
	"google.golang.org/grpc"
)

type server struct {
	listener net.Listener
	appointmentapi.UnimplementedAppointmentExternalServer
	appointmentapi.UnimplementedAppointmentInternalServer

	appointmentStore    AppointmentStore
	patientInfoService  patientinfoapi.PatientInfoServer
	notificationService notificationapi.NotificationServer
}

func newServer(store AppointmentStore) *server {
	return &server{
		appointmentStore: store,
	}
}

func (s *server) GRPCListenBlocking(port string) error {
	var err error
	s.listener, err = net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen at port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	appointmentapi.RegisterAppointmentExternalServer(grpcServer, newServer())
	appointmentapi.RegisterAppointmentInternalServer(grpcServer, newServer())

	log.Println("GRPPC server starts listening...")
	err = grpcServer.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (s *server) GetAppointmentsOnUser(c context.Context, in *appointmentapi.GetAppointmentsOnUserRequest) (*appointmentapi.GetAppointmentsReply, error) {
	// Validate input

	// Perform lookup
	internalAppointments, err := s.appointmentStore.GetAppointmentsOnUserUid(in.UserUid)
	if err != nil {
		return &appointmentapi.GetAppointmentsReply{
			Error: &appointmentapi.Error{
				Code:    500,
				Message: "Technical error fetching appointments on user",
				Details: err.Error(),
			},
		}, nil
	}

	return returnAppointments(internalAppointments), nil
}

func (s *server) GetAppointmentsOnStatus(c context.Context, in *appointmentapi.GetAppointmentsOnStatusRequest) (*appointmentapi.GetAppointmentsReply, error) {
	// Validate input

	// Perform lookup
	internalAppointments, err := s.appointmentStore.GetAppointmentsOnStatus(AppointmentStatus(in.Status))
	if err != nil {
		return &appointmentapi.GetAppointmentsReply{
			Error: &appointmentapi.Error{
				Code:    500,
				Message: "Technical error fetching appointments on user",
				Details: err.Error(),
			},
		}, nil
	}

	return returnAppointments(internalAppointments), nil
}

func (s *server) RequestAppointment(c context.Context, in *appointmentapi.RequestAppointmentRequest) (*appointmentapi.AppointmentReply, error) {
	// Validate input
	// TODO

	// Adjust datastore
	appointmentCreated, err := s.appointmentStore.PutAppointment(convertIntoInternal(*in.Appointment))
	if err != nil {
		return &appointmentapi.AppointmentReply{
			Error: &appointmentapi.Error{
				Code:    500,
				Message: "Technical error creating appointment",
				Details: err.Error(),
			},
		}, nil
	}
	return returnAppointment(appointmentCreated), nil
}

func (s *server) ModifyAppointmentStatus(c context.Context, in *appointmentapi.ModifyAppointmentStatusRequest) (*appointmentapi.AppointmentReply, error) {
	// Validate input
	// TODO

	// Perform lookup

	// Perform lookup
	internalAppointment, found, err := s.appointmentStore.GetAppointmentOnUid(in.AppointmentUid)
	if err != nil {
		return &appointmentapi.AppointmentReply{
			Error: &appointmentapi.Error{
				Code:    500,
				Message: "Technical error fetching appointments on user",
				Details: err.Error(),
			},
		}, nil
	}
	if !found {
		return &appointmentapi.AppointmentReply{
			Error: &appointmentapi.Error{
				Code:    404,
				Message: "Appointment with uid not found",
			},
		}, nil
	}
	// Fetch patient details
	resp, err := s.patientInfoService.GetPatientOnUid(c, &patientinfoapi.GetPatientOnUidRequest{PatientUid: internalAppointment.UserUID})
	if err != nil {
		return &appointmentapi.AppointmentReply{
			Error: &appointmentapi.Error{
				Code:    500,
				Message: "Technical error finding patient on uid",
				Details: err.Error(),
			},
		}, nil
	}
	if resp.Error != nil {
		return &appointmentapi.AppointmentReply{
			Error: &appointmentapi.Error{
				Code:    resp.Error.Code,
				Message: resp.Error.Message,
				Details: resp.Error.Details,
			},
		}, nil
	}

	// Send out sms
	_, _ = s.notificationService.SendSms(c, &notificationapi.SendSmsRequest{
		Sms: &notificationapi.SmsMessage{
			RecipientPhoneNumber: resp.Patient.PhoneNumber,
			Body:                 "Appointment confirmed", // TODO use template
		},
	})
	// TODO error checking

	// Send out email
	_, _ = s.notificationService.SendEmail(c, &notificationapi.SendEmailRequest{
		Email: &notificationapi.EmailMessage{
			RecipientEmailAddress: resp.Patient.PhoneNumber,
			Subject:               "Appointment confirmed",              // TODO use template
			Body:                  "Appointment confirmed with details", // TODO use template
		},
	})
	// TODO error checking

	// Adjust datastore
	internalAppointment.Status = AppointmentStatusConfirmed
	appointmentAdjusted, err := s.appointmentStore.PutAppointment(internalAppointment)
	return &appointmentapi.AppointmentReply{
		Error: &appointmentapi.Error{
			Code:    500,
			Message: "Error modifying apppointment",
			Details: resp.Error.Details,
		},
	}, nil

	return returnAppointment(appointmentAdjusted), nil
}

func returnAppointments(internalAppointments []Appointment) *appointmentapi.GetAppointmentsReply {

	response := &appointmentapi.GetAppointmentsReply{
		Appointments: []*appointmentapi.Appointment{},
	}
	for _, a := range internalAppointments {
		response.Appointments = append(response.Appointments, convertIntoExternal(a))
	}
	return response
}

func returnAppointment(internalAppointmnent Appointment) *appointmentapi.AppointmentReply {
	response := &appointmentapi.AppointmentReply{
		Appointment: convertIntoExternal(internalAppointmnent),
	}

	return response
}

func convertIntoExternal(a Appointment) *appointmentapi.Appointment {
	return &appointmentapi.Appointment{
		AppointmentUid: a.AppointmentUID,
		UserUid:        a.UserUID,
		DateTime:       a.DateTime,
		Location:       a.Location,
		Topic:          a.Topic,
		Status:         appointmentapi.AppointmentStatus(a.Status),
	}
}

func convertIntoInternal(a appointmentapi.Appointment) Appointment {
	return Appointment{
		AppointmentUID: a.AppointmentUid,
		UserUID:        a.UserUid,
		DateTime:       a.DateTime,
		Location:       a.Location,
		Topic:          a.Topic,
		Status:         AppointmentStatus(a.Status),
	}
}
