package main

import "fmt"

// START OMIT
type Datastorer interface { // ends with "er" // HL
	Put(key string, value interface{}) error
	Get(key string) (interface{}, bool, error)
	Remove(key string) error
}

type PatientService struct {
	datastore Datastorer // get a datastore injected via "constructor"
}

func main() {
	patientService := NewPatientService( NewSimpplisticDatastore())
	patientService.MarkAllergicToAntiBiotocs("patient-12345")
}

// END OMIT

type Patient struct {
	UID        string
	Allergies []string
}

func NewPatientService(datastore  Datastorer) *PatientService{
	return &PatientService{
		datastore:datastore,
	}
}

func (ps PatientService)MarkAllergicToAntiBiotocs(patientUID string) error {
	serialzedPatient, exists, err := ps.datastore.Get(patientUID)
	if err != nil {
		return fmt.Errorf("Technical error fetching patient:%s", err)
	}
	if !exists {
		return fmt.Errorf("Patient with uid %s does not exist",patientUID)
	}
	patient := serialzedPatient.(Patient)
	patient.Allergies  = append(patient.Allergies, "antibiotics")
	return ps.datastore.Put(patientUID, patient)
}

type SimpplisticDatastore struct {
	data map[string]interface{}
}

func NewSimpplisticDatastore() Datastorer {
	return &SimpplisticDatastore{
		data: map[string]interface{}{},
	}
}

func (ds *SimpplisticDatastore) Put(key string, value interface{}) error{
	ds.data[key] = value
	return nil
}

func (ds *SimpplisticDatastore) Get(key string) (interface{}, bool, error){
	value, found := ds.data[key]
	return value, found, nil
}

func (ds *SimpplisticDatastore) Remove(key string) error{
	delete(ds.data,key)
	return nil
}
