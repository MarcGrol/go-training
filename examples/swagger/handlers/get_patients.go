package handlers

import (
	"log"

	"github.com/MarcGrol/go-training/examples/swagger/service/models"
	"github.com/MarcGrol/go-training/examples/swagger/service/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

// GetPatients is the handler for the GET /patients endpoint
func (c *Controller) GetPatients(params operations.GetPatientsParams) middleware.Responder {
	patients, err := c.ps.GetAll()

	if err != nil {
		log.Println(err)

		return operations.NewGetPatientsDefault(500).WithPayload(&models.Error{
			Message: err.Error(),
		})
	}

	var swPatients []*models.Patient

	for _, patient := range patients {
		swPatients = append(swPatients, patient.ToSwaggerModel())
	}

	return operations.NewGetPatientsOK().WithPayload(swPatients)
}
