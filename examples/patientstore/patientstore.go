// Package patientstore allows for fetching and storing patients in a persistent store
package patientstore

import "fmt"

// PatientStorer is used to persist (create and modify) and fetch hospital patients
type PatientStorer interface {
	// GetOnUID fetches a patient based on its globally unique id
	// on success: a Patient is returned (and error is set to nil)
	// if the patient was not found, the second return parameter is set to false
	// on a technical error, the third parameter is not nil
	GetOnUID(uid string) (Patient, bool, error)

	// Store persists a patient
	// on success: the Patient is returned (and error is set to nil)
	// If the patient did not yet exist, a globally unique id is assigned to that patient before storing.
	// on a technical error, the third parameter is not nil
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
