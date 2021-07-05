package registrationService

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// START OMIT
func TestRegistrationSucces(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorer := NewMockPatientStorer(ctrl)
	mockPincoder := NewMockPincodeGenerator(ctrl)
	mockSmsSender := NewMockSmsSender(ctrl)

	patient := Patient{
		UID:         "123",
		Name:        "Marc",
		PhoneNumber: "31648928857",
	}

	mockStorer.EXPECT().PutPatientOnUid(patient).Return(nil)                // HL
	mockPincoder.EXPECT().GenerateRandomPin().Return(1234)                  // HL
	mockSmsSender.EXPECT().SendSms(fmt.Sprintf("+%s", patient.PhoneNumber), // HL
		"Finalize registration with pincode 1234").Return(nil) // HL

	sut := NewRegistrationService(mockStorer, mockPincoder, mockSmsSender)

	err := sut.RegisterPatient(patient)
	assert.NoError(t, err) // HL
}

// END OMIT
