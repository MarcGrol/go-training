package patientstore

// Patient holds all information of a hospital patient
type Patient struct {
	UID         string
	FullName    string
	AddressLine string
	Allergies   []string
}
