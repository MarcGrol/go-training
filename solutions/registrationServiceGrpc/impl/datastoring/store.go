package datastoring

import (
	"sync"

	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/datastorer"
)

type inMemoryPatientStore struct {
	sync.Mutex
	patients map[string]datastorer.Patient
}

func New() datastorer.PatientStorer {
	return &inMemoryPatientStore{
		patients: map[string]datastorer.Patient{},
	}
}
func (ps *inMemoryPatientStore) PutPatientOnUid(patient datastorer.Patient) error {
	ps.Lock()
	defer ps.Unlock()

	ps.patients[patient.UID] = patient

	return nil
}

func (ps *inMemoryPatientStore) GetPatientOnUid(patientUID string) (datastorer.Patient, bool, error) {
	ps.Lock()
	defer ps.Unlock()

	patient, found := ps.patients[patientUID]

	return patient, found, nil
}
