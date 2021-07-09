package datastorer

//go:generate mockgen -source=dataStorer.go -destination=dataStorerMocks.go -package=datastorer PatientStorer

type PatientStorer interface {
	GetPatientOnUid(uid string) (Patient, bool, error)
	PutPatientOnUid(patient Patient) error
}
