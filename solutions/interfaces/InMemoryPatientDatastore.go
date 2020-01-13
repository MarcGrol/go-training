package interfaces

var (
	_ Datastorer = &inMemoryPatientDatastore{}
)

func init() {
	New = newInMemoryPatientDatastore
}

type inMemoryPatientDatastore struct {
	data map[string]Patient
}

func newInMemoryPatientDatastore() Datastorer {
	return &inMemoryPatientDatastore{
		data: map[string]Patient{},
	}
}

func (ds *inMemoryPatientDatastore) Put(key string, patient Patient) error {
	ds.data[key] = patient
	return nil
}

func (ds *inMemoryPatientDatastore) Get(key string) (Patient, bool, error) {
	patient, found := ds.data[key]
	return patient, found, nil
}

func (ds *inMemoryPatientDatastore) Remove(key string) error {
	delete(ds.data, key)
	return nil
}
