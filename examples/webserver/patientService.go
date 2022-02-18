package main

import (
	"context"
	"fmt"
)

// START OMIT

type patientWebService struct {
	uuider       func() string
	patientStore PatientStore
}

func NewPatientService(uuider func() string, patientStore PatientStore) *patientWebService {
	return &patientWebService{
		uuider:       uuider,
		patientStore: patientStore,
	}
}

func (s *patientWebService) getPatientOnUID(c context.Context, patientUID string) (Patient, bool, error) {
	return s.patientStore.GetOnUid(c, patientUID)
}

// END OMIT

func (s *patientWebService) createPatient(c context.Context, patient Patient) (Patient, error) {
	patient.UID = s.uuider()

	return patient, s.patientStore.RunInTransaction(c, func(ctx context.Context) error {
		return s.patientStore.Create(c, patient)
	})
}

func (s *patientWebService) modifyPatientOnUid(c context.Context, patient Patient) error {
	return s.patientStore.RunInTransaction(c, func(ctx context.Context) error {
		_, found, err := s.getPatientOnUID(c, patient.UID)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("Not found")
		}

		return s.patientStore.Modify(c, patient)
	})
}

func (s *patientWebService) deletePatientOnUid(c context.Context, patientUID string) error {
	return s.patientStore.RunInTransaction(c, func(ctx context.Context) error {
		_, found, err := s.getPatientOnUID(c, patientUID)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("Not found")
		}
		return s.patientStore.Remove(c, patientUID)
	})
}
