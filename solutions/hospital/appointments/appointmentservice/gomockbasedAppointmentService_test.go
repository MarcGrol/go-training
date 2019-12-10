package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"

	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentapi"
	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentservice/appointmentstore"
	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"
	"github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"
)

func TestMockgenBasedGetAppointmentsOnUser(t *testing.T) {
	testCases := [...]struct {
		description            string
		createAppointmentStore func(ctlr *gomock.Controller) appointmentstore.AppointmentStore
		request                *appointmentapi.GetAppointmentsOnUserRequest
		expectedResponse       *appointmentapi.GetAppointmentsReply
	}{
		{
			description: "Invalid input: missing userUid",
			request:     &appointmentapi.GetAppointmentsOnUserRequest{},
			expectedResponse: &appointmentapi.GetAppointmentsReply{
				Error: &appointmentapi.Error{
					Code:    400,
					Message: "Invalid input",
					Details: "userUid",
				},
			},
		},
		{
			description: "Error fetching appointments on user",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentsOnUserUid(gomock.Any()).
					Return([]appointmentstore.Appointment{}, fmt.Errorf("Error fetching appointments on user"))
				return mock
			},
			request: &appointmentapi.GetAppointmentsOnUserRequest{UserUid: "myUserUid"},
			expectedResponse: &appointmentapi.GetAppointmentsReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Technical error fetching appointments on user",
					Details: "Error fetching appointments on user",
				},
			},
		},
		{
			description: "Success fetching appointments on user",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentsOnUserUid("myUserUid").
					Return([]appointmentstore.Appointment{exampleAppointment}, nil)
				return mock
			},
			request: &appointmentapi.GetAppointmentsOnUserRequest{UserUid: "myUserUid"},
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			var appointmentStore appointmentstore.AppointmentStore
			if tc.createAppointmentStore != nil {
				appointmentStore = tc.createAppointmentStore(ctrl)
			}
			service := newServer(appointmentStore, nil, nil)
			response, _ := service.GetAppointmentsOnUser(c, tc.request)
			t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
			if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestMockgenBasedRequestAppointment(t *testing.T) {
	testCases := [...]struct {
		description                string
		createAppointmentStore     func(ctlr *gomock.Controller) appointmentstore.AppointmentStore
		createPatientServiceClient func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient
		request                    *appointmentapi.RequestAppointmentRequest
		expectedResponse           *appointmentapi.AppointmentReply
	}{
		{
			description: "Invalid input: Missing appointment",
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: nil,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    400,
					Message: "Invalid input",
					Details: "appointment",
				},
			},
		},
		{
			description: "Invalid input: Missing userUid",
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: &appointmentapi.Appointment{ /* empty request */ },
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    400,
					Message: "Invalid input",
					Details: "userUid",
				},
			},
		},
		{
			description: "Invalid input: Missing dateTime",
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: &appointmentapi.Appointment{
					UserUid:  exampleAppointment.UserUID,
					DateTime: "", // should not be empty
					Location: exampleAppointment.Location,
					Topic:    exampleAppointment.Topic,
				},
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    400,
					Message: "Invalid input",
					Details: "dateTime",
				},
			},
		},
		{
			description: "Invalid input: Missing location",
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: &appointmentapi.Appointment{
					UserUid:  exampleAppointment.UserUID,
					DateTime: exampleAppointment.DateTime,
					Location: "", // should not be empty
					Topic:    exampleAppointment.Topic,
				},
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    400,
					Message: "Invalid input",
					Details: "location",
				},
			},
		},
		{
			description: "Invalid input: Missing topic",
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: &appointmentapi.Appointment{
					UserUid:  exampleAppointment.UserUID,
					DateTime: exampleAppointment.DateTime,
					Location: exampleAppointment.Location,
					Topic:    "", // should not be empty
				},
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    400,
					Message: "Invalid input",
					Details: "topic",
				},
			},
		},
		{
			description: "Error fetching patient",
			createPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Error: &patientinfoapi.Error{Code: 500, Message: "xxx", Details: "a"},
					}, nil)
				return mock
			},
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: &appointmentapi.Appointment{
					UserUid:  exampleAppointment.UserUID,
					DateTime: exampleAppointment.DateTime,
					Location: exampleAppointment.Location,
					Topic:    exampleAppointment.Topic,
				},
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Error fetching patient on uid",
					Details: "500: xxx (a)",
				},
			},
		},
		{
			description: "Error creating appointment",
			createPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil)
				return mock
			},
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().PutAppointment(gomock.Any()).
					Return(appointmentstore.Appointment{}, fmt.Errorf("yyy"))
				return mock
			},
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: &appointmentapi.Appointment{
					UserUid:  exampleAppointment.UserUID,
					DateTime: exampleAppointment.DateTime,
					Location: exampleAppointment.Location,
					Topic:    exampleAppointment.Topic,
				},
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Technical error creating appointment",
					Details: "yyy",
				},
			},
		},
		{
			description: "Success creating appointment",
			createPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil)
				return mock
			},
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().PutAppointment(gomock.Any()).
					Return(exampleAppointment, nil)
				return mock
			},
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var appointmentStore appointmentstore.AppointmentStore
			if tc.createAppointmentStore != nil {
				appointmentStore = tc.createAppointmentStore(ctrl)
			}
			var patientServiceClient patientinfoapi.PatientInfoClient
			if tc.createPatientServiceClient != nil {
				patientServiceClient = tc.createPatientServiceClient(ctrl)
			}
			service := newServer(appointmentStore, patientServiceClient, nil)
			response, _ := service.RequestAppointment(c, tc.request)
			t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
			if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestMockgenBasedGetAppointmentsOnStatus(t *testing.T) {
	testCases := [...]struct {
		description            string
		createAppointmentStore func(ctlr *gomock.Controller) appointmentstore.AppointmentStore
		request                *appointmentapi.GetAppointmentsOnStatusRequest
		expectedResponse       *appointmentapi.GetAppointmentsReply
	}{
		{
			description: "Invalid input: invalid status",
			request:     &appointmentapi.GetAppointmentsOnStatusRequest{Status: appointmentapi.AppointmentStatus_UNKNOWN},
			expectedResponse: &appointmentapi.GetAppointmentsReply{
				Error: &appointmentapi.Error{
					Code:    400,
					Message: "Invalid input",
					Details: "status",
				},
			},
		},
		{
			description: "Error fetching appointments on status",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentsOnStatus(gomock.Any()).
					Return([]appointmentstore.Appointment{}, fmt.Errorf("Error fetching appointments on status"))
				return mock
			},
			request: &appointmentapi.GetAppointmentsOnStatusRequest{Status: appointmentapi.AppointmentStatus_REQUESTED},
			expectedResponse: &appointmentapi.GetAppointmentsReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Technical error fetching appointments on status",
					Details: "Error fetching appointments on status",
				},
			},
		},
		{
			description: "Success fetching appointments on status",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentsOnStatus(gomock.Any()).
					Return([]appointmentstore.Appointment{exampleAppointment}, nil)
				return mock
			},
			request: &appointmentapi.GetAppointmentsOnStatusRequest{Status: appointmentapi.AppointmentStatus_REQUESTED},
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			var appointmentStore appointmentstore.AppointmentStore
			if tc.createAppointmentStore != nil {
				appointmentStore = tc.createAppointmentStore(ctrl)
			}
			service := newServer(appointmentStore, nil, nil)
			response, _ := service.GetAppointmentsOnStatus(c, tc.request)
			t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
			if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestMockgenBasedConfirmAppointment(t *testing.T) {
	testCases := [...]struct {
		description                     string
		createAppointmentStore          func(ctlr *gomock.Controller) appointmentstore.AppointmentStore
		createPatientServiceClient      func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient
		createNotificationServiceClient func(ctlr *gomock.Controller) notificationapi.NotificationClient
		request                         *appointmentapi.ModifyAppointmentStatusRequest
		expectedResponse                *appointmentapi.AppointmentReply
	}{
		{
			description: "Invalid input: appointmentUid",
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "", // should not be empty
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    400,
					Message: "Invalid input",
					Details: "appointmentUid",
				},
			},
		},
		{
			description: "Invalid input: status",
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_UNKNOWN, // should not be unknown
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    400,
					Message: "Invalid input",
					Details: "status",
				},
			},
		},
		{
			description: "Error fetching appointment",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid("myAppointmentUid").
					Return(appointmentstore.Appointment{}, false, fmt.Errorf("qqq"))
				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Error fetching appointment on uid",
					Details: "qqq",
				},
			},
		},
		{
			description: "Appointment not found",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid("myAppointmentUid").
					Return(appointmentstore.Appointment{}, false, nil)
				return mock
			},
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
			description: "Error fetching patient",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid("myAppointmentUid").
					Return(exampleAppointment, true, nil)
				return mock
			},
			createPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Error: &patientinfoapi.Error{Code: 500, Message: "yyy", Details: "d"},
					}, nil)
				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Error fetching patient on uid",
					Details: "500: yyy (d)",
				},
			}},
		{
			description: "Patient not found",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid("myAppointmentUid").
					Return(exampleAppointment, true, nil)
				return mock
			},
			createPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Error: &patientinfoapi.Error{Code: 404, Message: "xxx", Details: "a"},
					}, nil)
				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    404,
					Message: "Error fetching patient on uid",
					Details: "404: xxx (a)",
				},
			},
		},
		{
			description: "Error notifying via email",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid("myAppointmentUid").
					Return(exampleAppointment, true, nil)
				return mock
			},
			createPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil)
				return mock
			},
			createNotificationServiceClient: func(ctlr *gomock.Controller) notificationapi.NotificationClient {
				mock := notificationapi.NewMockNotificationClient(ctlr)
				mock.EXPECT().SendEmail(gomock.Any(), gomock.Any()).
					Return(&notificationapi.SendReply{
						Error: &notificationapi.Error{
							Code: 500, Message: "xxx", Details: "yyy",
						},
					}, nil)
				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Error sending email",
					Details: "500: xxx (yyy)",
				},
			},
		},
		{
			description: "Error notifying via sms",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid("myAppointmentUid").
					Return(exampleAppointment, true, nil)
				return mock
			},
			createPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil)
				return mock
			},
			createNotificationServiceClient: func(ctlr *gomock.Controller) notificationapi.NotificationClient {
				mock := notificationapi.NewMockNotificationClient(ctlr)
				mock.EXPECT().SendEmail(gomock.Any(), gomock.Any()).
					Return(&notificationapi.SendReply{
						Status: notificationapi.DeliveryStatus_DELIVERED}, nil)
				mock.EXPECT().SendSms(gomock.Any(), gomock.Any()).
					Return(&notificationapi.SendReply{
						Error: &notificationapi.Error{
							Code: 500, Message: "aaa", Details: "bbb",
						},
					}, nil)

				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Error sending sms",
					Details: "500: aaa (bbb)",
				},
			},
		},
		{
			description: "Error storing confirmed appointment",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid("myAppointmentUid").
					Return(exampleAppointment, true, nil)
				mock.EXPECT().PutAppointment(gomock.Any()).
					Return(appointmentstore.Appointment{}, fmt.Errorf("abc"))
				return mock
			},
			createPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil)
				return mock
			},
			createNotificationServiceClient: func(ctlr *gomock.Controller) notificationapi.NotificationClient {
				mock := notificationapi.NewMockNotificationClient(ctlr)
				mock.EXPECT().SendEmail(gomock.Any(), gomock.Any()).
					Return(&notificationapi.SendReply{
						Status: notificationapi.DeliveryStatus_DELIVERED}, nil)
				mock.EXPECT().SendSms(gomock.Any(), gomock.Any()).
					Return(&notificationapi.SendReply{
						Status: notificationapi.DeliveryStatus_DELIVERED}, nil)

				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: &appointmentapi.AppointmentReply{
				Error: &appointmentapi.Error{
					Code:    500,
					Message: "Error persisting modified appointment",
					Details: "abc",
				},
			},
		},
		{
			description: "Success confirming appointment",
			createAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid("myAppointmentUid").
					Return(exampleAppointment, true, nil)
				mock.EXPECT().PutAppointment(gomock.Any()).
					Return(exampleAppointment, nil)
				return mock
			},
			createPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil)
				return mock
			},
			createNotificationServiceClient: func(ctlr *gomock.Controller) notificationapi.NotificationClient {
				mock := notificationapi.NewMockNotificationClient(ctlr)
				mock.EXPECT().SendEmail(gomock.Any(), gomock.Any()).
					Return(&notificationapi.SendReply{
						Status: notificationapi.DeliveryStatus_DELIVERED}, nil)
				mock.EXPECT().SendSms(gomock.Any(), gomock.Any()).
					Return(&notificationapi.SendReply{
						Status: notificationapi.DeliveryStatus_DELIVERED}, nil)

				return mock
			},
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			var appointmentStore appointmentstore.AppointmentStore
			if tc.createAppointmentStore != nil {
				appointmentStore = tc.createAppointmentStore(ctrl)
			}
			var patientServiceClient patientinfoapi.PatientInfoClient
			if tc.createPatientServiceClient != nil {
				patientServiceClient = tc.createPatientServiceClient(ctrl)
			}
			var notificationServiceClient notificationapi.NotificationClient
			if tc.createNotificationServiceClient != nil {
				notificationServiceClient = tc.createNotificationServiceClient(ctrl)
			}
			service := newServer(appointmentStore, patientServiceClient, notificationServiceClient)
			response, _ := service.ModifyAppointmentStatus(c, tc.request)
			t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
			if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
				t.Error(diff)
			}
		})
	}
}
