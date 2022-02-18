package main

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/datastorer"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/emailsender"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/pincoder"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/smssender"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/uuider"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/regprotobuf"
)

type RegistrationService struct {
	uuidGenerator    uuider.UuidGenerator
	patientStore     datastorer.PatientStorer
	emailSender      emailsender.EmailSender
	smsSender        smssender.SmsSender
	pincodeGenerator pincoder.PincodeGenerator
	regprotobuf.UnimplementedRegistrationServiceServer
}

func NewRegistrationService(uuidGenerator uuider.UuidGenerator, patientStore datastorer.PatientStorer, pincoder pincoder.PincodeGenerator,
	emailSender emailsender.EmailSender, smsSender smssender.SmsSender) *RegistrationService {
	return &RegistrationService{
		uuidGenerator:    uuidGenerator,
		patientStore:     patientStore,
		pincodeGenerator: pincoder,
		emailSender:      emailSender,
		smsSender:        smsSender,
	}
}

func (rs *RegistrationService) RegisterPatient(ctx context.Context, req *regprotobuf.RegisterPatientRequest) (*regprotobuf.RegisterPatientResponse, error) {
	err := validateRegisterPatientRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error validating request: %s", err.Error())
	}

	pincode := rs.pincodeGenerator.GeneratePincode()
	if req.Patient.Contact.PhoneNumber != "" {
		smsContent := fmt.Sprintf("Finalize registration with pincode %d", pincode)
		err = rs.smsSender.SendSms(internationalize(req.Patient.Contact.PhoneNumber), smsContent)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Error sending sms: %s", err)
		}
	} else if req.Patient.Contact.EmailAddress != "" {
		emailSubject := "Registration pincode"
		emailContent := fmt.Sprintf("Finalize registration with pincode %d", pincode)
		err = rs.emailSender.SendEmail(req.Patient.Contact.EmailAddress, emailSubject, emailContent)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Error sending email: %s", err)
		}
	}

	patient := patientToInternal(req.Patient)
	patient.RegistrationPin = pincode
	patient.UID = rs.uuidGenerator.GenerateUuid()
	patient.RegistrationStatus = datastorer.Pending

	err = rs.patientStore.PutPatientOnUid(patient)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error storring patient: %s", err)
	}

	return &regprotobuf.RegisterPatientResponse{
		PatientUid: patient.UID,
	}, nil
}

func (rs *RegistrationService) CompletePatientRegistration(ctx context.Context, req *regprotobuf.CompletePatientRegistrationRequest) (*regprotobuf.CompletePatientRegistrationResponse, error) {
	err := validatePatientRegistrationRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Error validating input: %s", err)
	}

	patient, found, err := rs.patientStore.GetPatientOnUid(req.PatientUid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error getting patient in uid: %s", err)
	}
	if !found {
		return nil, status.Errorf(codes.NotFound, "Patient with uid not found")
	}

	if int(req.Credentials.Pincode) != patient.RegistrationPin {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid pin")
	}

	patient.RegistrationStatus = datastorer.Registered
	err = rs.patientStore.PutPatientOnUid(patient)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error storing patien: %s", err)
	}

	return &regprotobuf.CompletePatientRegistrationResponse{
		Status: regprotobuf.RegistrationStatus_REGISTRATION_CONFIRMED,
	}, nil
}

func validateRegisterPatientRequest(req *regprotobuf.RegisterPatientRequest) error {
	if req == nil || req.Patient == nil || req.Patient.BSN == "" || req.Patient.FullName == "" || req.Patient.Contact == nil {
		return fmt.Errorf("Missing base fields")
	}
	if req.Patient.Contact.PhoneNumber == "" && req.Patient.Contact.EmailAddress == "" {
		return fmt.Errorf("Missing contacts")
	}
	return nil
}

func validatePatientRegistrationRequest(req *regprotobuf.CompletePatientRegistrationRequest) error {
	if req == nil || req.PatientUid == "" || req.Credentials == nil || req.Credentials.Pincode == 0 {
		return fmt.Errorf("Missing credentials")
	}
	return nil
}

func internationalize(phoneNumber string) string {
	if strings.HasPrefix(phoneNumber, "+") {
		return phoneNumber
	}
	return "+" + phoneNumber
}

func patientToInternal(p *regprotobuf.Patient) datastorer.Patient {
	return datastorer.Patient{
		BSN:      p.BSN,
		FullName: p.FullName,
		Address: datastorer.StreetAddress{
			PostalCode: func() string {
				if p.Address != nil {
					return p.Address.PostalCode
				}
				return ""
			}(),
			HouseNumber: func() int {
				if p.Address != nil {
					return int(p.Address.HouseNumber)
				}
				return 0
			}(),
		},
		Contact: datastorer.Contact{
			PhoneNumber: func() string {
				if p.Contact != nil {
					return p.Contact.PhoneNumber
				}
				return ""
			}(),
			EmailAddress: func() string {
				if p.Contact != nil {
					return p.Contact.EmailAddress
				}
				return ""
			}(),
		},
	}

}
