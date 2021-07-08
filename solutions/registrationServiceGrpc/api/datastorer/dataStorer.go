package datastorer

import "github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/model"

//go:generate mockgen -source=dataStorer.go -destination=dataStorerMocks.go -package=datastorer PatientStorer

type PatientStorer interface {
	GetPatientOnUid(uid string) (model.Patient, bool, error)
	PutPatientOnUid(patient model.Patient) error
}
