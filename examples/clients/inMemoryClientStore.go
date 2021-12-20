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
	patients map[string]Client
}

func newClientStore(nower func() time.Time) ClientStore {
	return &inMemoryPatientStore{
		nower:    nower,
		patients: map[string]Client{},
	}
}

func (s *inMemoryPatientStore) Create(ctx context.Context, patient Client) error {
	s.Lock()
	defer s.Unlock()

	if patient.UID == "" {
		return fmt.Errorf("Invalid patient: missing uid")
	}
	patient.CreatedAt = s.nower()
	s.patients[patient.UID] = patient
	return nil
}

func (s *inMemoryPatientStore) Modify(ctx context.Context, patient Client) error {
	s.Lock()
	defer s.Unlock()

	if patient.UID == "" {
		return fmt.Errorf("Invalid patient: missing uid")
	}
	patient.LastModified = s.nower()
	s.patients[patient.UID] = patient
	return nil
}

func (s *inMemoryPatientStore) GetOnUid(ctx context.Context, appointmentUID string) (Client, bool, error) {
	s.Lock()
	defer s.Unlock()

	patient, found := s.patients[appointmentUID]
	return patient, found, nil
}

func (s *inMemoryPatientStore) Search(ctx context.Context) ([]Client, error) {
	s.Lock()
	defer s.Unlock()

	found := []Client{}
	for _, appointment := range s.patients {
		found = append(found, appointment)
	}
	return found, nil
}

func (s *inMemoryPatientStore) Remove(ctx context.Context, patientUID string) error {
	s.Lock()
	defer s.Unlock()

	delete(s.patients, patientUID)

	return nil
}
