package datastoring

import (
	datastorer2 "github.com/MarcGrol/go-training/examples/registrationServiceGrpc/api/datastorer"
	"sync"
)

type inMemoryPatientStore struct {
	sync.Mutex
	patients map[string]datastorer2.Patient
}

func New() datastorer2.PatientStorer {
	return &inMemoryPatientStore{
		patients: map[string]datastorer2.Patient{},
	}
}
func (ps *inMemoryPatientStore) PutPatientOnUid(patient datastorer2.Patient) error {
	ps.Lock()
	defer ps.Unlock()

	ps.patients[patient.UID] = patient

	return nil
}

func (ps *inMemoryPatientStore) GetPatientOnUid(patientUID string) (datastorer2.Patient, bool, error) {
	ps.Lock()
	defer ps.Unlock()

	patient, found := ps.patients[patientUID]

	return patient, found, nil
}
