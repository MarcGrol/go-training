
package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "github.com/MarcGrol/go-training/examples/project/notificationService/spec"
)


type Client struct {
	conn *grpc.ClientConn
	client pb.NotificationClient
}

func New() (*Client) {
	return &Client{}
}

func (c *Client)Open(addressPort string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addressPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("Error connecting: %v", err)
	}
	c.client = pb.NewNotificationClient(conn)
	return nil
}

func (c *Client)Close(){
	if c.conn != nil {
		defer c.conn.Close()
	}
}

func (c *Client)SendEmail(ctx context.Context, recipientEmailAddress, subject, body string) (string,error){
	response, err := c.client.SendEmail(ctx, &pb.SendEmailRequest{
		Email:&pb.EmailMessage{
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

func (c *Client)SendSms(ctx context.Context, recipientPhoneNumber, body string) (string,error){
	response, err := c.client.SendSms(ctx, &pb.SendSmsRequest{
		Sms:&pb.SmsMessage{
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

func (c *Client)GetStatus(ctx context.Context, msgUID string) (pb.NotificationStatus,error) {
	response, err := c.client.GetNotificationStatus(ctx, &pb.GetNotificationStatusRequest{
		NotificationUid: msgUID,
	})
	if err != nil {
		log.Fatalf("Error getting status: %v", err)
	}

	if response.Error != nil {
		return pb.NotificationStatus_UNKNOWN, fmt.Errorf("Error getting sms: %+v", response.Error)
	}
	return  response.GetStatus(), nil
}
