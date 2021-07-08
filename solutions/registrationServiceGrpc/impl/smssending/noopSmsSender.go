package smssending

import (
	"github.com/MarcGrol/go-training/solutions/registrationServiceGrpc/api/smssender"
)

type noopSmsSender struct{}

func New() smssender.SmsSender {
	return &noopSmsSender{}
}

func (ss *noopSmsSender) SendSms(phoneNumber string, smsContent string) error {
	return nil
}
