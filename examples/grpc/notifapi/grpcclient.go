
package notifapi

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type client struct {
	conn   *grpc.ClientConn
	client NotificationClient
}

func New() (Notifier) {
	return &client{}
}

func (c *client)Open(addressPort string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addressPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("Error connecting: %v", err)
	}
	c.client = NewNotificationClient(conn)
	return nil
}

func (c *client)Close(){
	if c.conn != nil {
		defer c.conn.Close()
	}
}

func (c *client)SendEmail(ctx context.Context, recipientEmailAddress, subject, body string) (string,error){
	response, err := c.client.SendEmail(ctx, &SendEmailRequest{
		Email:&EmailMessage{
			RecipientEmailAddress: recipientEmailAddress,
			Subject:               subject,
			Body:                  body,
		},
	})
	if err != nil{
		return "", fmt.Errorf("Could not send email: %v", err)
	}

	if response.Error != nil {
		return "", fmt.Errorf("Error sending email: %+v", response.Error)
	}

	return response.GetNotificationUid() , nil
}

func (c *client)SendSms(ctx context.Context, recipientPhoneNumber, body string) (string,error){
	response, err := c.client.SendSms(ctx, &SendSmsRequest{
		Sms:&SmsMessage{
			RecipientPhoneNumber: recipientPhoneNumber,
			Body:                 body,
		},
	})
	if err != nil {
		return "", fmt.Errorf("Could not send sms: %v", err)
	}

	if response.Error != nil {
		return "", fmt.Errorf("Error sending sms: %+v", response.Error)
	}
	return response.GetNotificationUid() , nil
}

func (c *client)GetStatus(ctx context.Context, msgUID string) (NotificationStatus,error) {
	response, err := c.client.GetNotificationStatus(ctx, &GetNotificationStatusRequest{
		NotificationUid: msgUID,
	})
	if err != nil {
		log.Fatalf("Error getting status: %v", err)
	}

	if response.Error != nil {
		return NotificationStatus_UNKNOWN, fmt.Errorf("Error getting sms: %+v", response.Error)
	}
	return  response.GetStatus(), nil
}
