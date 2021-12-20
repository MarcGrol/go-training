package main

import "log"

type SmsSender interface {
	SendSms(sender, recipient, body string) error
}

type dummySmsSender struct {

}


func NewDummySmsSender() SmsSender {
	return &dummySmsSender{}
}

func (s *dummySmsSender)SendSms(sender, recipient, body string) error {
	log.Printf("Simulate sending sms to '%s' with body '%s'", recipient, body)
	return nil
}