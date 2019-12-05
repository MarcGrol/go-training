package main

import "sync"

type Patient struct {
	UID          string
	FullName     string
	AddressLine  string
	PhoneNumber  string
	EmailAddress string
}

type PatientStore interface {
	GetPatientOnUid(uid string) (Patient, bool, error)
}

type patientStore struct {
	sync.Mutex // be prepared when this service starts supporting modifications
	patients   map[string]Patient
}

func newPatientStore() PatientStore {
	return &patientStore{
		patients: map[string]Patient{
			"1": {UID: "1", FullName: "Marc Grol", PhoneNumber: "+3148928856", EmailAddress: "mgrol@xebia.com"},
			"2": {UID: "2", FullName: "Eva Berkhout", PhoneNumber: "+31626656696", EmailAddress: "eva.marc@hetnet.nl"},
		},
	}
}

func (ps *patientStore) GetPatientOnUid(uid string) (Patient, bool, error) {
	ps.Lock()
	defer ps.Unlock()

	patient, found := ps.patients[uid]
	return patient, found, nil
}
