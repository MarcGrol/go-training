package main

import (
	"sync"
)

type Patient struct {
	UID          string    `json:"uid"`
	FullName     string    `json:"fullName"`
	AddressLine  string    `json:"addressLine"`
}

func newPatientStore(uider func() string) *inMemoryPatientStore {
	return &inMemoryPatientStore{
		patients: map[string]Patient{},
	}
}

// START OMIT
type inMemoryPatientStore struct {
	sync.Mutex
	patients map[string]Patient
}

func (ps *inMemoryPatientStore) Put(patient Patient) error {
	ps.Lock()
	defer ps.Unlock() // HL

	ps.patients[patient.UID] = patient

	return nil
}
// END OMIT
