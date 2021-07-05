package main

import (
	"context"
	"fmt"
	"time"
)

// START OMIT

type patientWebService struct {
	uuider       func() string
	nower        func() time.Time
	patientStore PatientStore
}

func NewPatientService(uuider func() string, nower func() time.Time, patientStore PatientStore) *patientWebService {
	return &patientWebService{
		uuider:       uuider,
		nower:        nower,
		patientStore: patientStore,
	}
}

func (s *patientWebService) getPatientOnUID(c context.Context, patientUID string) (Patient, bool, error) {
	return s.patientStore.GetOnUid(c, patientUID)
}

// END OMIT

func (s *patientWebService) createPatient(c context.Context, patient Patient) (Patient, error) {
	patient.UID = s.uuider()
	patient.CreatedAt = s.nower()
	return s.patientStore.Put(c, patient)
}

func (s *patientWebService) modifyPatientOnUid(c context.Context, patient Patient) (Patient, error) {
	_, found, err := s.getPatientOnUID(c, patient.UID)
	if err != nil {
		return Patient{}, err
	}
	if !found {
		return Patient{}, fmt.Errorf("Not found")
	}

	patient.LastModified = s.nower()

	return s.patientStore.Put(c, patient)
}

func (s *patientWebService) deletePatientOnUid(c context.Context, patientUID string) error {
	_, found, err := s.getPatientOnUID(c, patientUID)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("Not found")
	}
	return s.patientStore.Remove(c, patientUID)
}
