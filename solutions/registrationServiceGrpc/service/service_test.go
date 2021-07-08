package service

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/datastorer"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/model"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/pincoder"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/smssender"
)

// START OMIT
func TestRegistrationSucces(t *testing.T) {
	ctrl, mockStorer, mockPincoder, mockSmsSender := setupDependencies(t)
	defer ctrl.Finish()

	patient := model.Patient{
		UID:         "123",
		FullName:    "Marc",
		PhoneNumber: "31648928857",
	}

	mockStorer.EXPECT().PutPatientOnUid(gomock.Any()).Return(nil)           // HL
	mockPincoder.EXPECT().GeneratePincode().Return(1234)                    // HL
	mockSmsSender.EXPECT().SendSms(fmt.Sprintf("+%s", patient.PhoneNumber), // HL
		"Finalize registration with pincode 1234").Return(nil) // HL

	sut := NewRegistrationService(mockStorer, mockPincoder, mockSmsSender)

	err := sut.RegisterPatient(patient)
	assert.NoError(t, err) // HL
}

// END OMIT

func TestRegistrationWithoutPhoneNUmber(t *testing.T) {
	ctrl, mockStorer, mockPincoder, mockSmsSender := setupDependencies(t)
	defer ctrl.Finish()

	patient := model.Patient{
		UID:      "123",
		FullName: "Marc",
	}

	mockStorer.EXPECT().PutPatientOnUid(patient).Return(nil) // HL

	sut := NewRegistrationService(mockStorer, mockPincoder, mockSmsSender)

	err := sut.RegisterPatient(patient)
	assert.NoError(t, err) // HL
}

func TestRegistrationDatastoreError(t *testing.T) {
	ctrl, mockStorer, mockPincoder, mockSmsSender := setupDependencies(t)
	defer ctrl.Finish()

	patient := model.Patient{
		UID:      "123",
		FullName: "Marc",
	}

	mockStorer.EXPECT().PutPatientOnUid(patient).Return(fmt.Errorf("Store error")) // HL

	sut := NewRegistrationService(mockStorer, mockPincoder, mockSmsSender)

	err := sut.RegisterPatient(patient)
	assert.Error(t, err) // HL
}

func TestRegistrationDatastoreSmsSenderError(t *testing.T) {
	ctrl, mockStorer, mockPincoder, mockSmsSender := setupDependencies(t)
	defer ctrl.Finish()

	patient := model.Patient{
		UID:         "123",
		FullName:    "Marc",
		PhoneNumber: "31648928857",
	}

	mockStorer.EXPECT().PutPatientOnUid(patient).Return(nil)
	mockPincoder.EXPECT().GeneratePincode().Return(1234)
	mockSmsSender.EXPECT().SendSms(fmt.Sprintf("+%s", patient.PhoneNumber),
		"Finalize registration with pincode 1234").Return(fmt.Errorf("error contact remote service")) // HL

	sut := NewRegistrationService(mockStorer, mockPincoder, mockSmsSender)

	err := sut.RegisterPatient(patient)
	assert.Error(t, err) // HL
}

func setupDependencies(t *testing.T) (*gomock.Controller, *datastorer.MockPatientStorer, *pincoder.MockPincodeGenerator, *smssender.MockSmsSender) {
	ctrl := gomock.NewController(t)

	mockStorer := datastorer.NewMockPatientStorer(ctrl)
	mockPincoder := pincoder.NewMockPincodeGenerator(ctrl)
	mockSmsSender := smssender.NewMockSmsSender(ctrl)

	return ctrl, mockStorer, mockPincoder, mockSmsSender
}
