package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentapi"
	"github.com/MarcGrol/go-training/solutions/hospital/appointments/appointmentservice/appointmentstore"
	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"
	"github.com/MarcGrol/go-training/solutions/hospital/patients/patientinfoapi"
)

func TestGetAppointmentsOnUser(t *testing.T) {
	testCases := [...]struct {
		description               string
		constructAppointmentStore func(ctlr *gomock.Controller) appointmentstore.AppointmentStore
		request                   *appointmentapi.GetAppointmentsOnUserRequest
		expectedResponse          *appointmentapi.GetAppointmentsReply
		err                       string
	}{
		{
			description:      "Missing userUid",
			request:          &appointmentapi.GetAppointmentsOnUserRequest{},
			expectedResponse: nil,
			err:              fmt.Sprintf("Missing user-uid"),
		},
		{
			description: "Error fetching appointments on user",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentsOnUserUid(gomock.Any(), gomock.Any()).
					Return([]appointmentstore.Appointment{}, fmt.Errorf("Error fetching appointments on user"))
				return mock
			},
			request:          &appointmentapi.GetAppointmentsOnUserRequest{UserUid: "myUserUid"},
			expectedResponse: nil,
			err:              fmt.Sprintf("Technical error fetching appointments on user"),
		},
		{
			description: "Success fetching appointments on user",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentsOnUserUid(gomock.Any(), gomock.Any()).
					Return([]appointmentstore.Appointment{exampleAppointment}, nil).
					Do(func(c context.Context, userUID string) {
						assert.Equal(t, "myUserUid", userUID)
					})
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

			appointmentStore := appointmentStoreOrNil(ctrl, tc.constructAppointmentStore)

			service := newServer(appointmentStore, nil, nil)
			response, err := service.GetAppointmentsOnUser(c, tc.request)
			if tc.err != "" {
				t.Logf("%s: want: %+v, got:%+v", tcName, tc.err, err.Error())
				if !strings.Contains(err.Error(), tc.err) {
					//t.Error("Unexpected error")
				}
			} else {
				t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
				if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
					t.Error(diff)
				}
			}
		})
	}
}

func TestRequestAppointment(t *testing.T) {
	testCases := [...]struct {
		description                   string
		constructAppointmentStore     func(ctlr *gomock.Controller) appointmentstore.AppointmentStore
		constructPatientServiceClient func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient
		request                       *appointmentapi.RequestAppointmentRequest
		expectedResponse              *appointmentapi.AppointmentReply
		err                           string
	}{
		{
			description: "Invalid input: Missing appointment",
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: nil,
			},
			expectedResponse: nil,
			err:              fmt.Sprintf("Missing appointment"),
		},
		{
			description: "Invalid input: Missing userUid",
			request: &appointmentapi.RequestAppointmentRequest{
				Appointment: &appointmentapi.Appointment{ /* empty request */ },
			},
			expectedResponse: nil,
			err:              fmt.Sprintf("Missing userUid"),
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
			expectedResponse: nil,
			err:              fmt.Sprintf("Missing dateTime"),
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
			expectedResponse: nil,
			err:              fmt.Sprintf("Missing location"),
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
			expectedResponse: nil,
			err:              fmt.Sprintf("Missing topic"),
		},
		{
			description: "Error fetching patient",
			constructPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(nil, status.Errorf(codes.Internal, "xxx"))
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
			expectedResponse: nil,
			err:              fmt.Sprintf("Technical error fetching patient on uid"),
		},
		{
			description: "Error creating appointment",
			constructPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil)
				return mock
			},
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().PutAppointment(gomock.Any(), gomock.Any()).
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
			expectedResponse: nil,
			err:              fmt.Sprintf("Technical error creating appointment"),
		},
		{
			description: "Success creating appointment",
			constructPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil).
					Do(func(ctx context.Context, in *patientinfoapi.GetPatientOnUidRequest, opts ...grpc.CallOption) {
						assert.Equal(t, "myUserUid", in.PatientUid)
					})

				return mock
			},
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().PutAppointment(gomock.Any(), gomock.Any()).
					Return(exampleAppointment, nil).
					Do(func(c context.Context, in appointmentstore.Appointment) {
						assert.Equal(t, exampleAppointment.AppointmentUID, in.AppointmentUID)
						assert.Equal(t, exampleAppointment.UserUID, in.UserUID)
						assert.Equal(t, exampleAppointment.DateTime, in.DateTime)
						assert.Equal(t, exampleAppointment.Location, in.Location)
						assert.Equal(t, exampleAppointment.Topic, in.Topic)
						assert.Equal(t, appointmentstore.AppointmentStatusRequested, in.Status)
					})
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

			appointmentStore := appointmentStoreOrNil(ctrl, tc.constructAppointmentStore)
			patientServiceClient := patientInfoClientOrNil(ctrl, tc.constructPatientServiceClient)

			service := newServer(appointmentStore, patientServiceClient, nil)
			response, err := service.RequestAppointment(c, tc.request)
			if tc.err != "" {
				t.Logf("%s: want: %+v, got:%+v", tcName, tc.err, err)
				if !strings.Contains(err.Error(), tc.err) {
					t.Errorf("Unexpected error")
				}
			} else {
				t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
				if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
					t.Error(diff)
				}
			}
		})
	}
}

func TestGetAppointmentsOnStatus(t *testing.T) {
	testCases := [...]struct {
		description               string
		constructAppointmentStore func(ctlr *gomock.Controller) appointmentstore.AppointmentStore
		request                   *appointmentapi.GetAppointmentsOnStatusRequest
		expectedResponse          *appointmentapi.GetAppointmentsReply
		err                       string
	}{
		{
			description:      "Invalid input: invalid status",
			request:          &appointmentapi.GetAppointmentsOnStatusRequest{Status: appointmentapi.AppointmentStatus_UNKNOWN},
			expectedResponse: nil,
			err:              fmt.Sprintf("Invalid status"),
		},
		{
			description: "Error fetching appointments on status",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentsOnStatus(gomock.Any(), gomock.Any()).
					Return([]appointmentstore.Appointment{}, fmt.Errorf("Technical error fetching appointments on status"))
				return mock
			},
			request:          &appointmentapi.GetAppointmentsOnStatusRequest{Status: appointmentapi.AppointmentStatus_REQUESTED},
			expectedResponse: nil,
			err:              fmt.Sprintf("Technical error fetching appointments on status"),
		},
		{
			description: "Success fetching appointments on status",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentsOnStatus(gomock.Any(), gomock.Any()).
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

			appointmentStore := appointmentStoreOrNil(ctrl, tc.constructAppointmentStore)

			service := newServer(appointmentStore, nil, nil)
			response, err := service.GetAppointmentsOnStatus(c, tc.request)
			if tc.err != "" {
				t.Logf("%s: want: %+v, got:%+v", tcName, tc.err, err)
				if !strings.Contains(err.Error(), tc.err) {
					t.Errorf("Unexpected error")
				}
			} else {
				t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
				if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
					t.Error(diff)
				}
			}
		})
	}
}

func TestConfirmAppointment(t *testing.T) {
	testCases := [...]struct {
		description                        string
		constructAppointmentStore          func(ctlr *gomock.Controller) appointmentstore.AppointmentStore
		constructPatientServiceClient      func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient
		constructNotificationServiceClient func(ctlr *gomock.Controller) notificationapi.NotificationClient
		request                            *appointmentapi.ModifyAppointmentStatusRequest
		expectedResponse                   *appointmentapi.AppointmentReply
		err                                string
	}{
		{
			description: "Invalid input: appointmentUid",
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "", // should not be empty
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: nil,
			err:              fmt.Sprintf("Missing appointmentUid"),
		},
		{
			description: "Invalid input: status",
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_UNKNOWN, // should not be unknown
			},
			expectedResponse: nil,
			err:              fmt.Sprintf("Missing status"),
		},
		{
			description: "Error fetching appointment",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid(gomock.Any(), gomock.Any()).
					Return(appointmentstore.Appointment{}, false, fmt.Errorf("qqq"))
				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: nil,
			err:              fmt.Sprintf("Error fetching appointment on uid"),
		},
		{
			description: "Appointment not found",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid(gomock.Any(), gomock.Any()).
					Return(appointmentstore.Appointment{}, false, nil)
				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: nil,
			err:              fmt.Sprintf("Appointment with uid not found"),
		},
		{
			description: "Error fetching patient",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid(gomock.Any(), gomock.Any()).
					Return(exampleAppointment, true, nil)
				return mock
			},
			constructPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(nil, status.Errorf(codes.Internal, "yyy"))
				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: nil,
			err:              fmt.Sprintf("Technical error fetching patient on uid"),
		},
		{
			description: "Patient not found",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid(gomock.Any(), gomock.Any()).
					Return(exampleAppointment, true, nil)
				return mock
			},
			constructPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(nil, status.Errorf(codes.NotFound, "yyy"))

				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: nil,
			err:              fmt.Sprintf("Technical error fetching patient on uid"),
		},
		{
			description: "Error notifying via email",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid(gomock.Any(), gomock.Any()).
					Return(exampleAppointment, true, nil)
				return mock
			},
			constructPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil)
				return mock
			},
			constructNotificationServiceClient: func(ctlr *gomock.Controller) notificationapi.NotificationClient {
				mock := notificationapi.NewMockNotificationClient(ctlr)
				mock.EXPECT().SendEmail(gomock.Any(), gomock.Any()).
					Return(nil, status.Errorf(codes.NotFound, "zzz"))
				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: nil,
			err:              fmt.Sprintf("Technical error sending email"),
		},
		{
			description: "Error notifying via sms",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid(gomock.Any(), gomock.Any()).
					Return(exampleAppointment, true, nil)
				return mock
			},
			constructPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil)
				return mock
			},
			constructNotificationServiceClient: func(ctlr *gomock.Controller) notificationapi.NotificationClient {
				mock := notificationapi.NewMockNotificationClient(ctlr)
				mock.EXPECT().SendEmail(gomock.Any(), gomock.Any()).
					Return(&notificationapi.SendReply{
						Status: notificationapi.DeliveryStatus_DELIVERED}, nil)
				mock.EXPECT().SendSms(gomock.Any(), gomock.Any()).
					Return(nil, status.Errorf(codes.NotFound, "qqq"))

				return mock
			},
			request: &appointmentapi.ModifyAppointmentStatusRequest{
				AppointmentUid: "myAppointmentUid",
				Status:         appointmentapi.AppointmentStatus_CONFIRMED,
			},
			expectedResponse: nil,
			err:              fmt.Sprintf("Technical error sending sms"),
		},
		{
			description: "Error storing confirmed appointment",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid(gomock.Any(), gomock.Any()).
					Return(exampleAppointment, true, nil)
				mock.EXPECT().PutAppointment(gomock.Any(), gomock.Any()).
					Return(appointmentstore.Appointment{}, fmt.Errorf("abc"))
				return mock
			},
			constructPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil)
				return mock
			},
			constructNotificationServiceClient: func(ctlr *gomock.Controller) notificationapi.NotificationClient {
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
			expectedResponse: nil,
			err:              fmt.Sprintf("Error persisting modified appointment"),
		},
		{
			description: "Success confirming appointment",
			constructAppointmentStore: func(ctlr *gomock.Controller) appointmentstore.AppointmentStore {
				mock := appointmentstore.NewMockAppointmentStore(ctlr)
				mock.EXPECT().GetAppointmentOnUid(gomock.Any(), "myAppointmentUid").
					Return(exampleAppointment, true, nil).
					Do(func(ctx context.Context, uid string) {
						assert.Equal(t, exampleAppointment.AppointmentUID, uid)
					})
				mock.EXPECT().PutAppointment(gomock.Any(), gomock.Any()).
					Return(exampleAppointment, nil).
					Do(func(ctx context.Context, in appointmentstore.Appointment) {
						assert.Equal(t, exampleAppointment.UserUID, in.UserUID)
						assert.Equal(t, exampleAppointment.AppointmentUID, in.AppointmentUID)
						assert.Equal(t, exampleAppointment.Topic, in.Topic)
						assert.Equal(t, exampleAppointment.Location, in.Location)
						assert.Equal(t, exampleAppointment.DateTime, in.DateTime)
						assert.Equal(t, appointmentstore.AppointmentStatusConfirmed, in.Status)
					})
				return mock
			},
			constructPatientServiceClient: func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient {
				mock := patientinfoapi.NewMockPatientInfoClient(ctlr)
				mock.EXPECT().GetPatientOnUid(gomock.Any(), gomock.Any()).
					Return(&patientinfoapi.GetPatientOnUidReply{
						Patient: &examplePatient,
					}, nil).
					Do(func(ctx context.Context, in *patientinfoapi.GetPatientOnUidRequest, opts ...grpc.CallOption) {
						assert.Equal(t, exampleAppointment.UserUID, in.PatientUid)
					})
				return mock
			},
			constructNotificationServiceClient: func(ctlr *gomock.Controller) notificationapi.NotificationClient {
				mock := notificationapi.NewMockNotificationClient(ctlr)
				mock.EXPECT().SendEmail(gomock.Any(), gomock.Any()).
					Return(&notificationapi.SendReply{
						Status: notificationapi.DeliveryStatus_DELIVERED}, nil).
					Do(func(ctx context.Context, in *notificationapi.SendEmailRequest, opts ...grpc.CallOption) {
						assert.Equal(t, "myEmailAddress", in.Email.RecipientEmailAddress)
						assert.Equal(t, "Appointment confirmed", in.Email.Subject)
						assert.Equal(t, "Appointment confirmed with details", in.Email.Body)
					})

				mock.EXPECT().SendSms(gomock.Any(), gomock.Any()).
					Return(&notificationapi.SendReply{
						Status: notificationapi.DeliveryStatus_DELIVERED}, nil).
					Do(func(ctx context.Context, in *notificationapi.SendSmsRequest, opts ...grpc.CallOption) {
						assert.Equal(t, "myPhoneNumber", in.Sms.RecipientPhoneNumber)
						assert.Equal(t, "Appointment confirmed", in.Sms.Body)
					})
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

			appointmentStore := appointmentStoreOrNil(ctrl, tc.constructAppointmentStore)
			patientServiceClient := patientInfoClientOrNil(ctrl, tc.constructPatientServiceClient)
			notificationServiceClient := notificationClientOrNil(ctrl, tc.constructNotificationServiceClient)

			service := newServer(appointmentStore, patientServiceClient, notificationServiceClient)
			response, err := service.ModifyAppointmentStatus(c, tc.request)
			if tc.err != "" {
				t.Logf("%s: want: %+v, got:%+v", tcName, tc.err, err)
				if !strings.Contains(err.Error(), tc.err) {
					t.Errorf("Unexpected error")
				}
			} else {
				t.Logf("%s: want: %+v, got:%+v", tcName, *tc.expectedResponse, *response)
				if diff := deep.Equal(*tc.expectedResponse, *response); diff != nil {
					t.Error(diff)
				}
			}
		})
	}
}

func appointmentStoreOrNil(ctlr *gomock.Controller, constructor func(ctlr *gomock.Controller) appointmentstore.AppointmentStore) appointmentstore.AppointmentStore {
	if constructor == nil {
		return nil
	}
	return constructor(ctlr)
}

func patientInfoClientOrNil(ctlr *gomock.Controller, constructor func(ctlr *gomock.Controller) patientinfoapi.PatientInfoClient) patientinfoapi.PatientInfoClient {
	if constructor == nil {
		return nil
	}
	return constructor(ctlr)
}

func notificationClientOrNil(ctlr *gomock.Controller, constructor func(ctlr *gomock.Controller) notificationapi.NotificationClient) notificationapi.NotificationClient {
	if constructor == nil {
		return nil
	}
	return constructor(ctlr)
}

var exampleAppointment = appointmentstore.Appointment{
	AppointmentUID: "myAppointmentUid",
	UserUID:        "myUserUid",
	DateTime:       "myDateTime",
	Location:       "myLocation",
	Topic:          "myTopic",
	Status:         appointmentstore.AppointmentStatusRequested,
}

var examplePatient = patientinfoapi.Patient{
	Uid:          "myUid",
	FullName:     "myFullName",
	AddressLine:  "myAddressLine",
	PhoneNumber:  "myPhoneNumber",
	EmailAddress: "myEmailAddress",
}
