package main

import "log"

type EmailSender interface {
	SendEmail(sender, recipient, subject, body string) error
}


type dummyEmailSender struct {
}

func NewDummyEmailSender() EmailSender {
	return &dummyEmailSender{}
}

func (s *dummyEmailSender)SendEmail(sender, recipient, subject, body string) error {
	log.Printf("Simulate sending email to '%s' with subject '%s'", recipient, subject)
	return nil
}