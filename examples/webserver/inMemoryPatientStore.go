package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type inMemoryPatientStore struct {
	sync.Mutex
	nower    func() time.Time
	patients map[string]Patient
}

func newPatientStore(nower func() time.Time) PatientStore {
	return &inMemoryPatientStore{
		nower:    nower,
		patients: map[string]Patient{},
	}
}

func (as *inMemoryPatientStore) RunInTransaction(ctx context.Context, run withinTransactionFunc) error {
	as.Lock()
	defer as.Unlock()

	return run(ctx)
}

func (as *inMemoryPatientStore) Create(ctx context.Context, patient Patient) error {
	if patient.UID == "" {
		return fmt.Errorf("Invalid patient: missing uid")
	}
	patient.CreatedAt = as.nower()
	as.patients[patient.UID] = patient
	return nil
}

func (as *inMemoryPatientStore) Modify(ctx context.Context, patient Patient) error {
	if patient.UID == "" {
		return fmt.Errorf("Invalid patient: missing uid")
	}
	patient.LastModified = as.nower()
	as.patients[patient.UID] = patient
	return nil
}

func (as *inMemoryPatientStore) GetOnUid(ctx context.Context, appointmentUID string) (Patient, bool, error) {
	patient, found := as.patients[appointmentUID]
	return patient, found, nil
}

func (as *inMemoryPatientStore) Search(ctx context.Context) ([]Patient, error) {
	found := []Patient{}
	for _, appointment := range as.patients {
		found = append(found, appointment)
	}
	return found, nil
}

func (as *inMemoryPatientStore) Remove(ctx context.Context, patientUID string) error {
	delete(as.patients, patientUID)

	return nil
}
