package appointmentstore

import "github.com/google/uuid"

type Uider interface {
	Create() string
}

func NewBasicUuider() Uider {
	return &uuider{}
}

type uuider struct{}

func (u uuider) Create() string {
	return uuid.New().String()
}
