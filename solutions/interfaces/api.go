package interfaces

type Patient struct {
	UID  string
	Name string
}

type Datastorer interface {
	Put(key string, patient Patient) error
	Get(key string) (Patient, bool, error)
	Remove(key string) error
}

func New() Datastorer {
	return newInMemoryPatientDatastore()
}
