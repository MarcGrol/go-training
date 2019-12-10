package main

import (
	"context"
	"fmt"
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
		return &pb.GetAppointmentsReply{
			Error: convertApplicativeError(400, "Invalid input", "userUid"),
		}, nil
	}

	// Perform lookup
	internalAppointments, err := s.appointmentStore.GetAppointmentsOnUserUid(in.UserUid)
	if err != nil {
		return &pb.GetAppointmentsReply{
			Error: convertTechnicalError("Technical error fetching appointments on user", err),
		}, nil
	}

	return returnAppointmentList(internalAppointments), nil
}

func (s *server) GetAppointmentsOnStatus(c context.Context, in *pb.GetAppointmentsOnStatusRequest) (*pb.GetAppointmentsReply, error) {
	// Validate input
	if in == nil || in.GetStatus() == pb.AppointmentStatus_UNKNOWN {
		return &pb.GetAppointmentsReply{
			Error: convertApplicativeError(400, "Invalid input", "status"),
		}, nil
	}

	// Perform lookup
	internalAppointments, err := s.appointmentStore.GetAppointmentsOnStatus(appointmentstore.AppointmentStatus(in.Status))
	if err != nil {
		return &pb.GetAppointmentsReply{
			Error: convertTechnicalError("Technical error fetching appointments on status", err),
		}, nil
	}

	return returnAppointmentList(internalAppointments), nil
}

func (s *server) RequestAppointment(c context.Context, in *pb.RequestAppointmentRequest) (*pb.AppointmentReply, error) {
	// Validate input
	if in == nil || in.GetAppointment() == nil {
		return &pb.AppointmentReply{
			Error: convertApplicativeError(400, "Invalid input", "appointment"),
		}, nil
	} else {
		if in.GetAppointment().GetUserUid() == "" {
			return &pb.AppointmentReply{
				Error: convertApplicativeError(400, "Invalid input", "userUid"),
			}, nil
		}
		if in.GetAppointment().GetDateTime() == "" {
			return &pb.AppointmentReply{
				Error: convertApplicativeError(400, "Invalid input", "dateTime"),
			}, nil
		}
		if in.GetAppointment().GetLocation() == "" {
			return &pb.AppointmentReply{
				Error: convertApplicativeError(400, "Invalid input", "location"),
			}, nil
		}
		if in.GetAppointment().GetTopic() == "" {
			return &pb.AppointmentReply{
				Error: convertApplicativeError(400, "Invalid input", "topic"),
			}, nil
		}
	}

	// Check if patient exists
	getPatientResponse, err := s.patientInfoClient.GetPatientOnUid(c, &patientinfoapi.GetPatientOnUidRequest{PatientUid: in.Appointment.UserUid})
	if err != nil {
		return &pb.AppointmentReply{
			Error: convertTechnicalError("Technical error fetching patient on uid", err),
		}, nil
	}
	if getPatientResponse.Error != nil {
		return &pb.AppointmentReply{
			Error: convertApplicativeError(getPatientResponse.Error.Code, "Error fetching patient on uid",
				originalError(getPatientResponse.Error.Code, getPatientResponse.Error.Message, getPatientResponse.Error.Details)),
		}, nil
	}

	// Adjust datastore
	appointmentCreated, err := s.appointmentStore.PutAppointment(convertIntoInternal(*in.Appointment))
	if err != nil {
		return &pb.AppointmentReply{
			Error: convertTechnicalError("Technical error creating appointment", err),
		}, nil
	}
	log.Printf("Persisted new appointment")

	return returnSingleAppointment(appointmentCreated), nil
}

func (s *server) ModifyAppointmentStatus(c context.Context, in *pb.ModifyAppointmentStatusRequest) (*pb.AppointmentReply, error) {
	// Validate input
	// Validate input
	if in == nil {
		return &pb.AppointmentReply{
			Error: convertApplicativeError(400, "Invalid input", "request"),
		}, nil
	} else {
		if in.GetAppointmentUid() == "" {
			return &pb.AppointmentReply{
				Error: convertApplicativeError(400, "Invalid input", "appointmentUid"),
			}, nil
		}
		if in.GetStatus() == pb.AppointmentStatus_UNKNOWN {
			return &pb.AppointmentReply{
				Error: convertApplicativeError(400, "Invalid input", "status"),
			}, nil
		}
	}

	// Perform lookup
	internalAppointment, found, err := s.appointmentStore.GetAppointmentOnUid(in.AppointmentUid)
	if err != nil {
		return &pb.AppointmentReply{
			Error: convertTechnicalError("Error fetching appointment on uid", err),
		}, nil
	}
	if !found {
		return &pb.AppointmentReply{
			Error: convertApplicativeError(404, "Appointment with uid not found", ""),
		}, nil
	}
	log.Printf("Got appointment:%+v", internalAppointment)

	// Fetch patient details
	getPatientOnUidResp, err := s.patientInfoClient.GetPatientOnUid(c, &patientinfoapi.GetPatientOnUidRequest{PatientUid: internalAppointment.UserUID})
	if err != nil {
		return &pb.AppointmentReply{
			Error: convertTechnicalError("Technical error fetching patient on uid", err),
		}, nil
	}
	if getPatientOnUidResp.Error != nil {
		return &pb.AppointmentReply{
			Error: convertApplicativeError(getPatientOnUidResp.Error.Code, "Error fetching patient on uid",
				originalError(getPatientOnUidResp.Error.Code, getPatientOnUidResp.Error.Message, getPatientOnUidResp.Error.Details)),
		}, nil
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
			return &pb.AppointmentReply{
				Error: convertTechnicalError("Technical error sending email", err),
			}, nil
		}
		if sendEmailResponse.Error != nil {
			return &pb.AppointmentReply{
				Error: convertApplicativeError(sendEmailResponse.Error.Code, "Error sending email",
					originalError(sendEmailResponse.Error.Code, sendEmailResponse.Error.Message, sendEmailResponse.Error.Details)),
			}, nil
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
			return &pb.AppointmentReply{
				Error: convertTechnicalError("Technical error sending sms", err),
			}, nil
		}
		if sendSmsResponse.Error != nil {
			log.Printf("Error sending sms: %+v", sendSmsResponse.Error)
			return &pb.AppointmentReply{
				Error: convertApplicativeError(sendSmsResponse.Error.Code, "Error sending sms",
					originalError(sendSmsResponse.Error.Code, sendSmsResponse.Error.Message, sendSmsResponse.Error.Details)),
			}, nil
		}
		log.Printf("Send sms:%+v", sendSmsResponse)
	}

	// Adjust datastore
	internalAppointment.Status = appointmentstore.AppointmentStatusConfirmed
	appointmentAdjusted, err := s.appointmentStore.PutAppointment(internalAppointment)
	if err != nil {
		return &pb.AppointmentReply{
			Error: convertTechnicalError("Error persisting modified appointment", err),
		}, nil
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

func convertTechnicalError(errorMsg string, err error) *pb.Error {
	if err == nil {
		return nil
	}
	log.Printf("Got technical error: %s", err.Error())
	return &pb.Error{
		Code:    500,
		Message: errorMsg,
		Details: err.Error(),
	}
}

func convertApplicativeError(code int32, message, details string) *pb.Error {
	log.Printf("Got applicative error: %d: %s (%s)", code, message, details)
	return &pb.Error{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func originalError(code int32, message, details string) string {
	return fmt.Sprintf("%d: %s (%s)", code, message, details)
}
