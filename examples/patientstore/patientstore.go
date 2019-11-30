// Package patientstore allows for fetching and storing patients in a persistent store
package patientstore

import "fmt"

// PatientStorer is used to persist (create and modify) and fetch hospital patients
type PatientStorer interface {
	// GetOnUID fetches a patient based on its globally unique id
	// on a technical error, the error parameter is not nil
	// if the patient was not found, the second return parameter is set to false
	// on success: a Patient is returned
	GetOnUID(uid string) (Patient, bool, error)

	// Store persists a patient
	// on a technical error, the error parameter is not nil
	// on success: the Patient is returned
	// if the patient does not yet exist, UID (=globally unique) is set before pe
	Store(patient Patient) (Patient, error)
}

// New constructs a new patient-store
func New() PatientStorer {
	return &simplePatientStore{}
}

type simplePatientStore struct{}

func (ps *simplePatientStore) GetOnUID(uid string) (Patient, bool, error) {
	return Patient{}, false, fmt.Errorf("Not implemented")
}

func (ps *simplePatientStore) Store(patient Patient) (Patient, error) {
	return Patient{}, fmt.Errorf("Not implemented")
}
