package handlers

import "github.com/MarcGrol/go-training/examples/swagger/patientstore"

type Controller struct {
	ps patientstore.PatientStorer
}

func NewController(ps patientstore.PatientStorer) *Controller {
	return &Controller{
		ps,
	}
}
