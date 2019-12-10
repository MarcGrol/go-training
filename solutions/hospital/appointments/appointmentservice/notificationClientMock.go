package main

import (
	"context"

	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"

	"google.golang.org/grpc"
)

func NewNotificationClientMock(emailResponse *notificationapi.SendReply, smsResponse *notificationapi.SendReply) notificationapi.NotificationClient {
	return &mockNotificationClient{
		emailResponse: emailResponse,
		smsResponse:   smsResponse,
	}
}

type mockNotificationClient struct {
	emailResponse *notificationapi.SendReply
	smsResponse   *notificationapi.SendReply
}

func (m *mockNotificationClient) SendEmail(ctx context.Context, in *notificationapi.SendEmailRequest, opts ...grpc.CallOption) (*notificationapi.SendReply, error) {
	return m.emailResponse, nil

}
func (m *mockNotificationClient) SendSms(ctx context.Context, in *notificationapi.SendSmsRequest, opts ...grpc.CallOption) (*notificationapi.SendReply, error) {
	return m.smsResponse, nil
}
