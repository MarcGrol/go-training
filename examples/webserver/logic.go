package main

import (
	"fmt"
	"github.com/google/uuid"
)

type Patient struct {
	UID string
	FullName string
	AddressLine string
	Allergies []string
}

// START OMIT

type patientWebService struct {
}

func (s *patientWebService)getPatientOnUID(patientUID string) (Patient, error) {
	// Dummy implementation: a real service would use an inject datastore
	return Patient{
		UID:patientUID,
		FullName:"FirstName LastName",
		AddressLine:"Lindelaan 13, Groenekan",
		Allergies:[]string{"pinda", "antibiotics"},
	}, nil
}
// END OMIT

func (s *patientWebService)createPatient(patient Patient) (Patient, error) {
	patient.UID = uuid.New().String()
	return patient, nil
}

func (s *patientWebService)modifyPatientOnUid(patient Patient) (Patient, error) {
	return Patient{}, fmt.Errorf("Not implemented")
}

func (s *patientWebService)deletePatientOnUid(patient Patient) (error) {
	return fmt.Errorf("Not implemented")
}


