package main

import (
	"context"

	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"

	"google.golang.org/grpc"
)

func NewNotificationClientMock(emailResponse *notificationapi.SendEmailReply, smsResponse *notificationapi.SendSmsReply) notificationapi.NotificationClient {
	return &mockNotificationClient{
		emailResponse: emailResponse,
		smsResponse:   smsResponse,
	}
}

type mockNotificationClient struct {
	emailResponse *notificationapi.SendEmailReply
	smsResponse   *notificationapi.SendSmsReply
}

func (m *mockNotificationClient) SendEmail(ctx context.Context, in *notificationapi.SendEmailRequest, opts ...grpc.CallOption) (*notificationapi.SendEmailReply, error) {
	return m.emailResponse, nil

}
func (m *mockNotificationClient) SendSms(ctx context.Context, in *notificationapi.SendSmsRequest, opts ...grpc.CallOption) (*notificationapi.SendSmsReply, error) {
	return m.smsResponse, nil
}
