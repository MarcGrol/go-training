package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/uuider"

	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/datastorer"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/model"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/pincoder"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/smssender"
	pb "github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/impl/registration"
)

type RegistrationService struct {
	uuider           uuider.UuidGenerator
	patientStore     datastorer.PatientStorer
	smsSender        smssender.SmsSender
	pincodeGenerator pincoder.PincodeGenerator
}

// START OMIT
func NewRegistrationService(patientStore datastorer.PatientStorer, pincoder pincoder.PincodeGenerator, smsSender smssender.SmsSender) pb.RegistrationServiceServer {
	return &RegistrationService{
		patientStore: patientStore, pincodeGenerator: pincoder, smsSender: smsSender,
	}
}

func (rs *RegistrationService) RegisterPatient(ctx context.Context, req *pb.RegisterPatientRequest) (*pb.RegisterPatientResponse, error) {
	err := validateRegisterPatientRequest(req)
	if err != nil {
		return nil, err
	}

	patient := fillPatient(req.Patient)

	patient.UID = rs.uuider.GenerateUuid()
	patient.RegistrationPin = rs.pincodeGenerator.GeneratePincode()

	err = rs.patientStore.PutPatientOnUid(patient)
	if err != nil {
		return nil, err
	}

	if patient.Contact.PhoneNumber != "" {
		patient.RegistrationStatus = model.Pending
		pincode := rs.pincodeGenerator.GeneratePincode() // HL
		smsContent := fmt.Sprintf("Finalize registration with pincode %d", pincode)

		err = rs.smsSender.SendSms(internationalize(patient.Contact.PhoneNumber), smsContent) // HL
		if err != nil {
			return nil, err
		}
	}
	return &pb.RegisterPatientResponse{
		PatientUid: patient.UID,
	}, nil
}

func (rs *RegistrationService) CompletePatientRegistration(ctx context.Context, req *pb.CompletePatientRegistrationRequest) (*pb.CompletePatientRegistrationResponse, error) {
	err := validateCompletePatientRegistrationRequest(req)
	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	patient, found, err := rs.patientStore.GetPatientOnUid(req.Credentials.PatientUid)
	if err != nil {
		return nil, fmt.Errorf("Error getting patient in uid")
	}
	if !found {
		return nil, fmt.Errorf("Patient with uid not found")
	}

	if patient.RegistrationStatus == model.Locked {
		return nil, fmt.Errorf("Locked")
	}

	patient.PinRegistationAttempts++

	if int(req.Credentials.Pincode) != patient.RegistrationPin {
		if patient.PinRegistationAttempts > 5 {
			patient.RegistrationStatus = model.Locked
		}
		rs.patientStore.PutPatientOnUid(patient)
		return nil, fmt.Errorf("Invalid pin")
	}

	patient.PinRegistationAttempts = 0
	patient.RegistrationStatus = model.Registered

	err = rs.patientStore.PutPatientOnUid(patient)
	if err != nil {
		return nil, err
	}

	return &pb.CompletePatientRegistrationResponse{
		Status: pb.RegistrationStatus_REGISTRATION_CONFIRMED,
	}, nil
}

func validateRegisterPatientRequest(req *pb.RegisterPatientRequest) error {
	if req == nil || req.Patient == nil || req.Patient.BSN == "" || req.Patient.FullName == "" {
		return fmt.Errorf("Invalid patient")
	}
	return nil
}

func validateCompletePatientRegistrationRequest(req *pb.CompletePatientRegistrationRequest) error {
	if req == nil || req.Credentials == nil || req.Credentials.PatientUid == "" {
		return fmt.Errorf("Invalid credentials")
	}
	return nil
}

func internationalize(phoneNumber string) string {
	if strings.HasPrefix(phoneNumber, "+") {
		return phoneNumber
	}
	return "+" + phoneNumber
}

func fillPatient(p *pb.Patient) model.Patient {
	return model.Patient{
		BSN:      p.BSN,
		FullName: p.FullName,
		Address: model.StreetAddress{
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
		Contact: model.Contact{
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
