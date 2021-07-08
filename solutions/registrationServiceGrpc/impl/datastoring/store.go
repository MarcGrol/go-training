package datastoring

import (
	"sync"

	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/datastorer"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/model"
)

type inMemoryPatientStore struct {
	sync.Mutex
	patients map[string]model.Patient
}

func New() datastorer.PatientStorer {
	return &inMemoryPatientStore{
		patients: map[string]model.Patient{},
	}
}
func (ps *inMemoryPatientStore) PutPatientOnUid(patient model.Patient) error {
	ps.Lock()
	defer ps.Unlock()

	ps.patients[patient.UID] = patient

	return nil
}

func (ps *inMemoryPatientStore) GetPatientOnUid(patientUID string) (model.Patient, bool, error) {
	ps.Lock()
	defer ps.Unlock()

	patient, found := ps.patients[patientUID]

	return patient, found, nil
}
