package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/datastorer"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/emailsender"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/pincoder"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/smssender"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/uuider"
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/regprotobuf"
)

func TestRegistrationWithPhoneNumber(t *testing.T) {
	ctrl, uuidGenerator, mockStorer, mockPincoder, emailsender, mockSmsSender := setupDependencies(t)
	defer ctrl.Finish()

	req := &regprotobuf.RegisterPatientRequest{
		Patient: &regprotobuf.Patient{
			BSN:      "123",
			FullName: "Marc",
			Contact: &regprotobuf.Contact{
				PhoneNumber: "31648928857",
			},
		},
	}

	mockPincoder.EXPECT().GeneratePincode().Times(1).Return(1234)
	mockSmsSender.EXPECT().SendSms(fmt.Sprintf("+%s", req.Patient.Contact.PhoneNumber),
		"Finalize registration with pincode 1234").Return(nil)
	uuidGenerator.EXPECT().GenerateUuid().Return("abc123")
	mockStorer.EXPECT().PutPatientOnUid(gomock.Any()).Return(nil)

	sut := NewRegistrationService(uuidGenerator, mockStorer, mockPincoder, emailsender, mockSmsSender)

	resp, err := sut.RegisterPatient(context.TODO(), req)
	assert.NoError(t, err)
	assert.Equal(t, "abc123", resp.PatientUid)
}

func TestRegistrationWithEmail(t *testing.T) {
	ctrl, uuidGenerator, mockStorer, mockPincoder, emailsender, mockSmsSender := setupDependencies(t)
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

	sut := NewRegistrationService(uuidGenerator, mockStorer, mockPincoder, emailsender, mockSmsSender)

	resp, err := sut.RegisterPatient(context.TODO(), req)
	assert.NoError(t, err)
	assert.Equal(t, "abc123", resp.PatientUid)
}

func TestRegistrationInvalidInput(t *testing.T) {
	ctrl, uuidGenerator, mockStorer, mockPincoder, emailsender, mockSmsSender := setupDependencies(t)
	defer ctrl.Finish()

	req := &regprotobuf.RegisterPatientRequest{
		Patient: &regprotobuf.Patient{
			BSN: "123",
		},
	}
	sut := NewRegistrationService(uuidGenerator, mockStorer, mockPincoder, emailsender, mockSmsSender)

	resp, err := sut.RegisterPatient(context.TODO(), req)
	assert.Error(t, err) // internal error
	assert.Nil(t, resp)
}

func TestRegistrationDatastoreError(t *testing.T) {
	ctrl, uuidGenerator, mockStorer, mockPincoder, emailsender, mockSmsSender := setupDependencies(t)
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
	mockStorer.EXPECT().PutPatientOnUid(gomock.Any()).Return(fmt.Errorf("Store error"))

	sut := NewRegistrationService(uuidGenerator, mockStorer, mockPincoder, emailsender, mockSmsSender)

	resp, err := sut.RegisterPatient(context.TODO(), req)
	assert.Error(t, err) // internal error
	assert.Nil(t, resp)
}

func TestRegistrationDatastoreSmsSenderError(t *testing.T) {
	ctrl, uuidGenerator, mockStorer, mockPincoder, emailsender, mockSmsSender := setupDependencies(t)
	defer ctrl.Finish()

	req := &regprotobuf.RegisterPatientRequest{
		Patient: &regprotobuf.Patient{
			BSN:      "123",
			FullName: "Marc",
			Contact: &regprotobuf.Contact{
				PhoneNumber: "31648928857",
			},
		},
	}

	mockPincoder.EXPECT().GeneratePincode().Return(1234)
	mockSmsSender.EXPECT().SendSms(fmt.Sprintf("+%s", req.Patient.Contact.PhoneNumber),
		"Finalize registration with pincode 1234").Return(fmt.Errorf("error contact remote service"))

	sut := NewRegistrationService(uuidGenerator, mockStorer, mockPincoder, emailsender, mockSmsSender)

	resp, err := sut.RegisterPatient(context.TODO(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func setupDependencies(t *testing.T) (*gomock.Controller, *uuider.MockUuidGenerator, *datastorer.MockPatientStorer,
	*pincoder.MockPincodeGenerator, *emailsender.MockEmailSender, *smssender.MockSmsSender) {
	ctrl := gomock.NewController(t)

	uuidGenerator := uuider.NewMockUuidGenerator(ctrl)
	mockStorer := datastorer.NewMockPatientStorer(ctrl)
	mockPincoder := pincoder.NewMockPincodeGenerator(ctrl)
	mockEmailSender := emailsender.NewMockEmailSender(ctrl)
	mockSmsSender := smssender.NewMockSmsSender(ctrl)

	return ctrl, uuidGenerator, mockStorer, mockPincoder, mockEmailSender, mockSmsSender
}
