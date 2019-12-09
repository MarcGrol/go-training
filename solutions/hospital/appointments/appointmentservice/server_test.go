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

	c := context.TODO()
	for idx, tc := range testCases {
		tcName := fmt.Sprintf("Testcase: %d (%s)", idx, tc.description)
		t.Run(tcName, func(t *testing.T) {
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

	c := context.TODO()
	for idx, tc := range testCases {
		tcName := fmt.Sprintf("Testcase: %d (%s)", idx, tc.description)
		t.Run(tcName, func(t *testing.T) {
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

	c := context.TODO()
	for idx, tc := range testCases {
		tcName := fmt.Sprintf("Testcase: %d (%s)", idx, tc.description)
		t.Run(fmt.Sprintf("Testcase: %d", idx), func(t *testing.T) {
			service := newServer(tc.appointmentStore, nil, nil)
			response, _ := service.GetAppointmentsOnStatus(c, tc.request)
			t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
			if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestConfirmAppointment(t *testing.T) {
	testCases := []struct {
		description        string
		appointmentStore   AppointmentStore
		patientService     patientinfoapi.PatientInfoClient
		notificationClient notificationapi.NotificationClient
		request            *appointmentapi.ModifyAppointmentStatusRequest
		expectedResponse   *appointmentapi.AppointmentReply
	}{
		{
			description:        "Error fetching appointment",
			appointmentStore:   NewErrorMockAppointmentStore(errors.New("c")),
			patientService:     nil,
			notificationClient: nil,
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Error getting appointment on uid",
					Details: "c",
				},
			},
		},
		{
			description:        "Appointment not found",
			appointmentStore:   NewNotFoundMockAppointmentStore(),
			patientService:     nil,
			notificationClient: nil,
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    404,
					Message: "Appointment with uid not found",
				},
			},
		},
		{
			description:      "Error fetching patient",
			appointmentStore: NewsSuccesMockAppointmentStore(),
			patientService: NewPatientClientMock(&patientinfoapi.GetPatientOnUidReply{
				Error: &patientinfoapi.Error{
					Code:    500,
					Message: "yyy",
					Details: "d",
				},
			}),
			notificationClient: nil,
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Error getting patient on uid:yyy",
					Details: "d",
				},
			}},
		{
			description:      "Patient not found",
			appointmentStore: NewsSuccesMockAppointmentStore(),
			patientService: NewPatientClientMock(&patientinfoapi.GetPatientOnUidReply{
				Error: &patientinfoapi.Error{
					Code:    404,
					Message: "yyy",
				},
			}),
			notificationClient: nil,
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    404,
					Message: "Error getting patient on uid:yyy",
				},
			},
		},
		{
			description:      "Error notifying email",
			appointmentStore: NewsSuccesMockAppointmentStore(),
			patientService: NewPatientClientMock(&patientinfoapi.GetPatientOnUidReply{
				Patient: &examplePatient,
			}),
			notificationClient: NewNotificationClientMock(
				&notificationapi.SendEmailReply{
					Error: &notificationapi.Error{
						Code:    500,
						Message: "sss",
						Details: "ttt",
					},
				},
				nil,
			),
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "sss",
					Details: "ttt",
				},
			},
		},
		{
			description:      "Error notifying sms",
			appointmentStore: NewsSuccesMockAppointmentStore(),
			patientService: NewPatientClientMock(&patientinfoapi.GetPatientOnUidReply{
				Patient: &examplePatient,
			}),
			notificationClient: NewNotificationClientMock(
				&notificationapi.SendEmailReply{
					Status: notificationapi.DeliveryStatus_DELIVERED,
				},
				&notificationapi.SendSmsReply{
					Error: &notificationapi.Error{
						Code:    500,
						Message: "xxx",
						Details: "yyy",
					},
				},
			),
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "xxx",
					Details: "yyy",
				},
			},
		},
		//{
		//	description: "Error storing confirmed appointment",
		//	// TODO
		//},
		{
			description:      "Success confirming appointment",
			appointmentStore: NewsSuccesMockAppointmentStore(),
			patientService: NewPatientClientMock(&patientinfoapi.GetPatientOnUidReply{
				Patient: &examplePatient,
			}),
			notificationClient: NewNotificationClientMock(
				&notificationapi.SendEmailReply{
					Status: notificationapi.DeliveryStatus_DELIVERED,
				},
				&notificationapi.SendSmsReply{
					Status: notificationapi.DeliveryStatus_DELIVERED,
				},
			),
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
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

	c := context.TODO()
	for idx, tc := range testCases {
		tcName := fmt.Sprintf("Testcase: %d (%s)", idx, tc.description)
		t.Run(fmt.Sprintf("Testcase: %d", idx), func(t *testing.T) {
			service := newServer(tc.appointmentStore, tc.patientService, tc.notificationClient)
			response, _ := service.ModifyAppointmentStatus(c, tc.request)
			t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
			if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
				t.Error(diff)
			}
		})
	}
}
