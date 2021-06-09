package patientstore

import "github.com/MarcGrol/go-training/examples/swagger/service/models"

// Patient holds all information of a hospital patient
type Patient struct {
	ID              string
	FullName        string
	TelephoneNumber string
	Email           string
}

func (p *Patient) ToSwaggerModel() *models.Patient {
	return &models.Patient{
		Email:           &p.Email,
		FullName:        &p.FullName,
		ID:              &p.ID,
		TelephoneNumber: &p.TelephoneNumber,
	}
}
