package main

import (
	"log"

	"github.com/MarcGrol/go-training/solutions/hospital/notifications/notificationapi"
)

type fakeEmailSender struct{}

func (es fakeEmailSender) Send(recipientEmail, subject, emailBody string) error {
	log.Printf("Fake send email to %s: subject: %s, body: %s", recipientEmail, subject, emailBody)
	return nil
}

type fakeSmsSender struct{}

func (ss fakeSmsSender) Send(recipientPhoneNumber, messageBody string) error {
	log.Printf("Fake send sms to %s: %s", recipientPhoneNumber, messageBody)
	return nil
}

func main() {
	s := newServer(fakeEmailSender{}, fakeSmsSender{})
	err := s.GRPCListenBlocking(notificationapi.DefaultPort)
	if err != nil {
		log.Fatalf("Error starting rest-notification server: %s", err)
	}
}
