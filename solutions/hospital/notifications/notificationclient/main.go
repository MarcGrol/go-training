package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s mode recipent message", os.Args[0])
	}
	mode := os.Args[1]
	recipient := os.Args[2]
	message := os.Args[3]

	client, cleanup, err := notificationapi.NewGrpcClient(notificationapi.DefaultPort)
	if err != nil {
		log.Fatalf("*** Error creating motification-client: %v", err)
	}
	defer cleanup()
	log.Printf("Created motification-client")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if mode == "sms" {
		resp, err := client.SendSms(ctx, &notificationapi.SendSmsRequest{
			Sms: &notificationapi.SmsMessage{
				RecipientPhoneNumber: recipient,
				Body:                 message,
			},
		})
		if err != nil {
			log.Fatalf("sending sms: %s", err)
		}
		log.Printf("Sms delivery status: %s", resp.Status)

	} else if mode == "email" {
		resp, err := client.SendEmail(ctx, &notificationapi.SendEmailRequest{
			Email: &notificationapi.EmailMessage{
				RecipientEmailAddress: recipient,
				Subject:               message,
				Body:                  message,
			},
		})
		if err != nil {
			log.Fatalf("sending email: %s", err)
		}
		log.Printf("Email delivery status: %s", resp.Status)
	} else {
		log.Fatalf("Unknown mode %s", mode)
	}
}
