package main

import (
	"fmt"
)

type Patient struct {
	UID       string
	FullName  string
	Allergies []string
}

type PatientService struct {
	datastore Datastorer
}

// Constructor-like function: "injects" DataStorer
func NewPatientService(datastore Datastorer) *PatientService {
	ds := &PatientService{
		datastore: datastore,
	}
	return ds
}

func (ps PatientService) Create(patient Patient) error {
	err := ps.datastore.Put(patient.UID, patient)
	if err != nil {
		return fmt.Errorf("Technical error creating patient with uid:%s", err)
	}
	return nil
}

func (ps PatientService) MarkAllergicToAntiBiotics(patientUID string) error {
	opaque, exists, err := ps.datastore.Get(patientUID)
	if err != nil {
		return fmt.Errorf("Technical error fetching patient: %s", err)
	}
	if !exists {
		return fmt.Errorf("Patient with uid %s does not exist", patientUID)
	}
	patient := opaque.(Patient)
	patient.Allergies = append(patient.Allergies, "antibiotics")
	err = ps.datastore.Put(patientUID, patient)
	if err != nil {
		return fmt.Errorf("Technical error updating patient allergies: %s", err)
	}
	return nil
}
