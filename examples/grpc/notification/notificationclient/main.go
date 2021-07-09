package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MarcGrol/go-training/examples/grpc/notification/notificationapi"
)

const (
	address = "localhost:50051"
	reps    = 10
)

func main() {
	client, cleanup, err := notificationapi.NewGrpcClient(address)
	if err != nil {
		log.Fatalf("*** Error ccreating client: %v", err)
	}
	defer cleanup()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for idx := 0; idx < reps; idx++ {
		uid, err := sendEmail(ctx, client, "mgrol@xebia.com", fmt.Sprintf("My subject %d", idx), "My body")
		if err != nil {
			log.Printf("*** Error sending email: %v", err)
		} else {
			log.Printf("SendEmail-response: %+v", uid)
			status, err := getStatus(ctx, client, uid)
			if err != nil {
				log.Printf("*** Error getting status of sms: %v", err)
			} else {
				log.Printf("GetStatus-response on email with uid %s: %+v", uid, status)
			}
		}
	}

	for idx := 0; idx < reps; idx++ {
		uid, err := sendSms(ctx, client, "+31648928856", fmt.Sprintf("My body %d", idx))
		if err != nil {
			log.Printf("*** Error sending sms: %v", err)
		} else {
			log.Printf("SendEmail-response: %+v", uid)

			status, err := getStatus(ctx, client, uid)
			if err != nil {
				log.Printf("*** Error getting status of sms: %v", err)
			} else {
				log.Printf("GetStatus-response on sms with uid %s: %+v", uid, status)
			}
		}
	}
}

func sendEmail(ctx context.Context, c notificationapi.NotificationClient, recipientEmailAddress, subject, body string) (string, error) {
	response, err := c.SendEmail(ctx, &notificationapi.SendEmailRequest{
		Email: &notificationapi.EmailMessage{
			RecipientEmailAddress: recipientEmailAddress,
			Subject:               subject,
			Body:                  body,
		},
	})
	if err != nil {
		return "", fmt.Errorf("Could not send email: %v", err)
	}

	if response.Error != nil {
		return "", fmt.Errorf("Error sending email: %+v", response.Error)
	}

	return response.GetNotificationUid(), nil
}

func sendSms(ctx context.Context, c notificationapi.NotificationClient, recipientPhoneNumber, body string) (string, error) {
	response, err := c.SendSms(ctx, &notificationapi.SendSmsRequest{
		Sms: &notificationapi.SmsMessage{
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
	return response.GetNotificationUid(), nil
}

func getStatus(ctx context.Context, c notificationapi.NotificationClient, msgUID string) (notificationapi.NotificationStatus, error) {
	response, err := c.GetNotificationStatus(ctx, &notificationapi.GetNotificationStatusRequest{
		NotificationUid: msgUID,
	})
	if err != nil {
		log.Fatalf("Error getting status: %v", err)
	}

	if response.Error != nil {
		return notificationapi.NotificationStatus_UNKNOWN, fmt.Errorf("Error getting sms: %+v", response.Error)
	}
	return response.GetStatus(), nil
}
