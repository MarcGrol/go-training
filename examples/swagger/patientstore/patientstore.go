package patientstore

import (
	"errors"

	"github.com/google/uuid"
)

// PatientStorer is used to persist (create and modify) and fetch hospital patients
type PatientStorer interface {
	GetByID(id string) (*Patient, error)
	Store(patient Patient) (*Patient, error)
	GetAll() ([]*Patient, error)
}

// New constructs a new patient-store
func New() PatientStorer {
	return &simplePatientStore{
		patients: make(map[string]*Patient, 0),
	}
}

type simplePatientStore struct {
	patients map[string]*Patient
}

func (ps *simplePatientStore) GetByID(id string) (*Patient, error) {
	if ps.patients[id] == nil {
		return nil, errors.New("not found")
	}
	return ps.patients[id], nil
}

func (ps *simplePatientStore) Store(patient Patient) (*Patient, error) {
	patient.ID = uuid.New().String()
	ps.patients[patient.ID] = &patient

	return ps.patients[patient.ID], nil
}

func (ps *simplePatientStore) GetAll() ([]*Patient, error) {
	var patients []*Patient

	for _, patient := range ps.patients {
		patients = append(patients, patient)
	}

	return patients, nil
}
