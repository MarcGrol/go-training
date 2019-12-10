package main

import (
	"context"
	"sync"
)

type PatientStore interface {
	Put(ctx context.Context, appointment Patient) (Patient, error)
	GetOnUid(ctx context.Context, appointmentUID string) (Patient, bool, error)
	Search(ctx context.Context) ([]Patient, error)
	Remove(ctx context.Context, userUID string) error
}

type inMemoryPatientStore struct {
	sync.Mutex
	uider    Uider
	patients map[string]Patient
}

func newPatientStore(uider Uider) PatientStore {
	return &inMemoryPatientStore{
		uider:    uider,
		patients: map[string]Patient{},
	}
}

func (as *inMemoryPatientStore) Put(ctx context.Context, patient Patient) (Patient, error) {
	as.Lock()
	defer as.Unlock()

	if patient.UID == "" {
		patient.UID = as.uider.Create()
	}
	as.patients[patient.UID] = patient
	return patient, nil
}

func (as *inMemoryPatientStore) GetOnUid(ctx context.Context, appointmentUID string) (Patient, bool, error) {
	as.Lock()
	defer as.Unlock()

	patient, found := as.patients[appointmentUID]
	return patient, found, nil
}

func (as *inMemoryPatientStore) Search(ctx context.Context) ([]Patient, error) {
	as.Lock()
	defer as.Unlock()

	found := []Patient{}
	for _, appointment := range as.patients {
		found = append(found, appointment)
	}
	return found, nil
}

func (as *inMemoryPatientStore) Remove(ctx context.Context, patientUID string) error {
	as.Lock()
	defer as.Unlock()

	delete(as.patients, patientUID)

	return nil
}
