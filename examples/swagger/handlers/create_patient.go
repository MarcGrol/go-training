package handlers

import (
	"log"

	"github.com/MarcGrol/go-training/examples/swagger/service/models"

	"github.com/MarcGrol/go-training/examples/swagger/patientstore"
	"github.com/MarcGrol/go-training/examples/swagger/service/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// CreatePatient is the handler for the POST /patients endpoint
func (c *Controller) CreatePatient(params operations.CreatePatientParams) middleware.Responder {
	patient, err := c.ps.Store(patientstore.Patient{
		FullName:        *params.Data.FullName,
		TelephoneNumber: *params.Data.TelephoneNumber,
		Email:           params.Data.Email.String(),
	})

	if err != nil {
		log.Println(err)

		return operations.NewCreatePatientDefault(500).WithPayload(&models.Error{
			Message: err.Error(),
		})
	}

	return operations.NewCreatePatientOK().WithPayload(patient.ToSwaggerModel())
}
