package main

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/go-test/deep"

	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentapi"
	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"
	"github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"
)

func TestGetAppointmentsOnUser(t *testing.T) {
	testCases := [...]struct {
		description      string
		appointmentStore AppointmentStore
		request          *appointmentapi.GetAppointmentsOnUserRequest
		expectedResponse *appointmentapi.GetAppointmentsReply
	}{
		{
			description:      "Error fetching appointments on user",
			appointmentStore: NewErrorMockAppointmentStore(errors.New("a")),
			request:          &appointmentapi.GetAppointmentsOnUserRequest{UserUid: "myUserid"},
			expectedResponse: &appointmentapi.GetAppointmentsReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Technical error fetching appointments on user",
					Details: "a",
				},
			},
		},
		{
			description:      "Success fetching appointments on user",
			appointmentStore: NewsSuccesMockAppointmentStore(),
			request:          &appointmentapi.GetAppointmentsOnUserRequest{UserUid: "myuserid"},
			expectedResponse: &appointmentapi.GetAppointmentsReply{
				Appointments: []*appointmentapi.Appointment{
					{
						AppointmentUid: exampleAppointment.AppointmentUID,
						UserUid:        exampleAppointment.UserUID,
						DateTime:       exampleAppointment.DateTime,
						Location:       exampleAppointment.Location,
						Topic:          exampleAppointment.Topic,
						Status:         appointmentapi.AppointmentStatus(exampleAppointment.Status),
					},
				},
			},
		},
	}
	for idx, tc := range testCases {
		tcName := fmt.Sprintf("Testcase: %d (%s)", idx, tc.description)
		t.Run(tcName, func(t *testing.T) {
			c := context.TODO()
			service := newServer(tc.appointmentStore, nil, nil)
			response, _ := service.GetAppointmentsOnUser(c, tc.request)
			t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
			if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestRequestAppointment(t *testing.T) {
	testCases := []struct {
		description      string
		appointmentStore AppointmentStore
		patientService   patientinfoapi.PatientInfoClient
		request          *appointmentapi.RequestAppointmentRequest
		expectedResponse *appointmentapi.AppointmentReply
	}{
		{
			description: "Error fetching patient",
			patientService: NewPatientClientMock(&patientinfoapi.GetPatientOnUidReply{
				Error: &patientinfoapi.Error{
					Code:    500,
					Message: "xxx",
					Details: "a",
				},
			}),
			appointmentStore: nil,
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: &appointmentapi.Appointment{
					UserUid: exampleAppointment.UserUID,
				},
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Error getting patient on uid:xxx",
					Details: "a",
				},
			},
		},
		{
			description: "Error creating appointment",
			patientService: NewPatientClientMock(&patientinfoapi.GetPatientOnUidReply{
				Patient: &examplePatient,
			}),
			appointmentStore: NewErrorMockAppointmentStore(errors.New("b")),
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: &appointmentapi.Appointment{
					UserUid: exampleAppointment.UserUID,
				},
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Technical error creating appointment",
					Details: "b",
				},
			},
		},
		{
			description: "Success creating appointment",
			patientService: NewPatientClientMock(&patientinfoapi.GetPatientOnUidReply{
				Patient: &examplePatient,
			}),
			appointmentStore: NewsSuccesMockAppointmentStore(),
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: &appointmentapi.Appointment{
					AppointmentUid: exampleAppointment.AppointmentUID,
					UserUid:        exampleAppointment.UserUID,
					DateTime:       exampleAppointment.DateTime,
					Location:       exampleAppointment.Location,
					Topic:          exampleAppointment.Topic,
					Status:         appointmentapi.AppointmentStatus(exampleAppointment.Status),
				},
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Appointment: &appointmentapi.Appointment{
					AppointmentUid: exampleAppointment.AppointmentUID,
					UserUid:        exampleAppointment.UserUID,
					DateTime:       exampleAppointment.DateTime,
					Location:       exampleAppointment.Location,
					Topic:          exampleAppointment.Topic,
					Status:         appointmentapi.AppointmentStatus(exampleAppointment.Status),
				},
			},
		},
	}
	for idx, tc := range testCases {
		tcName := fmt.Sprintf("Testcase: %d (%s)", idx, tc.description)
		t.Run(tcName, func(t *testing.T) {
			c := context.TODO()
			service := newServer(tc.appointmentStore, tc.patientService, nil)
			response, _ := service.RequestAppointment(c, tc.request)
			t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
			if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestGetAppointmentsOnStatus(t *testing.T) {
	testCases := []struct {
		description      string
		appointmentStore AppointmentStore
		request          *appointmentapi.GetAppointmentsOnStatusRequest
		expectedResponse *appointmentapi.GetAppointmentsReply
	}{
		{
			description:      "Error fetching appointments on status",
			appointmentStore: NewErrorMockAppointmentStore(errors.New("a")),
			request:          &appointmentapi.GetAppointmentsOnStatusRequest{Status: appointmentapi.AppointmentStatus_REQUESTED},
			expectedResponse: &appointmentapi.GetAppointmentsReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Technical error fetching appointments on user",
					Details: "a",
				},
			},
		},
		{
			description:      "Success fetching appointments on status",
			appointmentStore: NewsSuccesMockAppointmentStore(),
			request:          &appointmentapi.GetAppointmentsOnStatusRequest{Status: appointmentapi.AppointmentStatus_REQUESTED},
			expectedResponse: &appointmentapi.GetAppointmentsReply{
				Appointments: []*appointmentapi.Appointment{
					{
						AppointmentUid: exampleAppointment.AppointmentUID,
						UserUid:        exampleAppointment.UserUID,
						DateTime:       exampleAppointment.DateTime,
						Location:       exampleAppointment.Location,
						Topic:          exampleAppointment.Topic,
						Status:         appointmentapi.AppointmentStatus(exampleAppointment.Status),
					},
				},
			},
		},
	}
	for idx, tc := range testCases {
		tcName := fmt.Sprintf("Testcase: %d (%s)", idx, tc.description)
		t.Run(fmt.Sprintf("Testcase: %d", idx), func(t *testing.T) {
			c := context.TODO()
			service := newServer(tc.appointmentStore, nil, nil)
			response, _ := service.GetAppointmentsOnStatus(c, tc.request)
			t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
			if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TesConfirmAppointment(t *testing.T) {
	testCases := []struct {
		description        string
		appointmentStore   AppointmentStore
		patientService     patientinfoapi.PatientInfoClient
		notificationClient notificationapi.NotificationClient
		request            *appointmentapi.ModifyAppointmentStatusRequest
		expectedResponse   *appointmentapi.AppointmentReply
	}{
		{
			description: "Error fetching appointment",
		},
		{
			description: "Appointment not found",
		},
		{
			description: "Error fetching patient",
		},
		{
			description: "Patient not found",
		},
		{
			description: "Error notifying email",
		},
		{
			description: "Error notifying sms",
		},
		{
			description: "Error storing confirmed appointment",
		},
		{
			description: "Success confirming appointment",
		},
	}
	for idx, tc := range testCases {
		tcName := fmt.Sprintf("Testcase: %d (%s)", idx, tc.description)
		t.Run(fmt.Sprintf("Testcase: %d", idx), func(t *testing.T) {
			c := context.TODO()
			service := newServer(tc.appointmentStore, tc.patientService, tc.notificationClient)
			response, _ := service.ModifyAppointmentStatus(c, tc.request)
			t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
			if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
				t.Error(diff)
			}
		})
	}
}
