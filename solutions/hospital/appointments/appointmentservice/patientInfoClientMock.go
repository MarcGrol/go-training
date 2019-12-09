package main

import (
	"context"

	"google.golang.org/grpc"

	"github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"
)

func NewPatientClientMock(response *patientinfoapi.GetPatientOnUidReply) patientinfoapi.PatientInfoClient {
	return &mockPatientInfoClient{
		response: response,
	}
}

type mockPatientInfoClient struct {
	response *patientinfoapi.GetPatientOnUidReply
}

func (m *mockPatientInfoClient) GetPatientOnUid(ctx context.Context, in *patientinfoapi.GetPatientOnUidRequest, opts ...grpc.CallOption) (*patientinfoapi.GetPatientOnUidReply, error) {
	return m.response, nil
}

var examplePatient = patientinfoapi.Patient{
	Uid:          "myUid",
	FullName:     "myFullName",
	AddressLine:  "myAddressLine",
	PhoneNumber:  "myPhoneNumber",
	EmailAddress: "myEmailAddress",
}
