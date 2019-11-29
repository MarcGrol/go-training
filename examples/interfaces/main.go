package main

import "log"

//START OMIT
func main() {
	// Inject Storer into business logic service
	patientService := NewPatientService(NewSimpplisticDatastore())

	patient := Patient{UID: "patient-12345", FullName: "Sjoerd Sjoerdsma", Allergies: []string{"pinda"}}

	// Initialize with data
	err := patientService.Create(patient) // uses Datastorer.Put // HL
	if err != nil {
		log.Fatalf("Error creating patient: %s", err)
	}

	// Adjust patient
	err = patientService.MarkAllergicToAntiBiotics(patient.UID) // uses Datastorer.Get and Put // HL
	if err != nil {
		log.Fatalf("MarkAllergicToAntiBiotics error: %s", err)
	}
}

//END OMIT
