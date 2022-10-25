package interfaces

type inMemoryPatientDatastore struct {
	data map[string]Patient
}

func newInMemoryPatientDatastore() Datastorer {
	return &inMemoryPatientDatastore{
		data: map[string]Patient{},
	}
}

func (ds *inMemoryPatientDatastore) Put(key string, patient any) error {
	ds.data[key] = patient.(Patient)
	return nil
}

func (ds *inMemoryPatientDatastore) Get(key string) (any, bool, error) {
	patient, found := ds.data[key]
	return patient, found, nil
}

func (ds *inMemoryPatientDatastore) Remove(key string) error {
	delete(ds.data, key)
	return nil
}
