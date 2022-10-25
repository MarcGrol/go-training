package interfaces

type Patient struct {
	UID  string
	Name string
}

type Datastorer interface {
	Put(key string, patient any) error
	Get(key string) (any, bool, error)
	Remove(key string) error
}

func New() Datastorer {
	return newInMemoryPatientDatastore()
}
