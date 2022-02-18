package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/api/datastorer"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/api/emailsender"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/api/pincoder"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/api/uuider"
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/regprotobuf"

)


func TestRegistrationWithEmail(t *testing.T) {
	ctrl, uuidGenerator, mockStorer, mockPincoder, emailsender := setupDependencies(t)
	defer ctrl.Finish()

	req := &regprotobuf.RegisterPatientRequest{
		Patient: &regprotobuf.Patient{
			BSN:      "123",
			FullName: "Marc",
			Contact: &regprotobuf.Contact{
				EmailAddress: "me@home.nl",
			},
		},
	}

	mockPincoder.EXPECT().GeneratePincode().Return(1234)
	emailsender.EXPECT().SendEmail(req.Patient.Contact.EmailAddress,
		"Registration pincode",
		"Finalize registration with pincode 1234").Return(nil)
	uuidGenerator.EXPECT().GenerateUuid().Return("abc123")
	mockStorer.EXPECT().PutPatientOnUid(gomock.Any()).Return(nil)

	sut := NewRegistrationService(uuidGenerator, mockStorer, mockPincoder, emailsender)

	resp, err := sut.RegisterPatient(context.TODO(), req)
	assert.NoError(t, err)
	assert.Equal(t, "abc123", resp.PatientUid)
}

func TestRegistrationInvalidInput(t *testing.T) {
	ctrl, uuidGenerator, mockStorer, mockPincoder, emailsender := setupDependencies(t)
	defer ctrl.Finish()

	req := &regprotobuf.RegisterPatientRequest{
		Patient: &regprotobuf.Patient{
			BSN: "123",
		},
	}
	sut := NewRegistrationService(uuidGenerator, mockStorer, mockPincoder, emailsender)

	resp, err := sut.RegisterPatient(context.TODO(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestRegistrationDatastoreError(t *testing.T) {
	ctrl, uuidGenerator, mockStorer, mockPincoder, emailsender := setupDependencies(t)
	defer ctrl.Finish()

	req := &regprotobuf.RegisterPatientRequest{
		Patient: &regprotobuf.Patient{
			BSN:      "123",
			FullName: "Marc",
			Contact: &regprotobuf.Contact{
				EmailAddress: "me@home.nl",
			},
		},
	}

	mockPincoder.EXPECT().GeneratePincode().Return(1234)
	emailsender.EXPECT().SendEmail(req.Patient.Contact.EmailAddress,
		gomock.Any(), gomock.Any()).Return(nil)
	uuidGenerator.EXPECT().GenerateUuid().Return("abc123")
	mockStorer.EXPECT().PutPatientOnUid(gomock.Any()).Return(fmt.Errorf("Store error"))

	sut := NewRegistrationService(uuidGenerator, mockStorer, mockPincoder, emailsender)

	resp, err := sut.RegisterPatient(context.TODO(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestRegistrationDatastoreEmailSenderError(t *testing.T) {
	ctrl, uuidGenerator, mockStorer, mockPincoder, mockEmailsender := setupDependencies(t)
	defer ctrl.Finish()

	req := &regprotobuf.RegisterPatientRequest{
		Patient: &regprotobuf.Patient{
			BSN:      "123",
			FullName: "Marc",
			Contact: &regprotobuf.Contact{
				EmailAddress: "me@home",
			},
		},
	}

	mockPincoder.EXPECT().GeneratePincode().Return(1234)
	mockEmailsender.EXPECT().SendEmail(req.Patient.Contact.EmailAddress,
		gomock.Any(), gomock.Any()).Return(fmt.Errorf("error contact remote service"))

	sut := NewRegistrationService(uuidGenerator, mockStorer, mockPincoder, mockEmailsender)

	resp, err := sut.RegisterPatient(context.TODO(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func setupDependencies(t *testing.T) (*gomock.Controller, *uuider.MockUuidGenerator, *datastorer.MockPatientStorer,
	*pincoder.MockPincodeGenerator, *emailsender.MockEmailSender) {
	ctrl := gomock.NewController(t)

	uuidGenerator := uuider.NewMockUuidGenerator(ctrl)
	mockStorer := datastorer.NewMockPatientStorer(ctrl)
	mockPincoder := pincoder.NewMockPincodeGenerator(ctrl)
	mockEmailSender := emailsender.NewMockEmailSender(ctrl)

	return ctrl, uuidGenerator, mockStorer, mockPincoder, mockEmailSender
}
