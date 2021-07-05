package registrationService

import (
	"fmt"
	"strings"
)

type RegistrationService struct {
	patientStore     PatientStorer
	smsSender        SmsSender
	pincodeGenerator PincodeGenerator
}

// START OMIT
func NewRegistrationService(patientStore PatientStorer, pincoder PincodeGenerator, smsSender SmsSender) *RegistrationService {
	return &RegistrationService{
		patientStore: patientStore, pincodeGenerator: pincoder, smsSender: smsSender,
	}
}

func (rs *RegistrationService) RegisterPatient(patient Patient) error {
	err := rs.patientStore.PutPatientOnUid(patient) // HL
	if err != nil {
		return err
	}

	if patient.PhoneNumber != "" {
		pincode := rs.pincodeGenerator.GenerateRandomPin() // HL
		smsContent := fmt.Sprintf("Finalize registration with pincode %d", pincode)

		err = rs.smsSender.SendSms(internationalize(patient.PhoneNumber), smsContent) // HL
		if err != nil {
			return err
		}
	}
	return nil
}

// END OMIT

func validate(p Patient) error {
	if p.UID == "" || p.Name == "" {
		return fmt.Errorf("Invalid patient")
	}
	return nil
}

func internationalize(phoneNumber string) string {
	if strings.HasPrefix(phoneNumber, "+") {
		return phoneNumber
	}
	return "+" + phoneNumber
}
