package smssending

import (
	"context"
)

// SmsClient is used to send SMSs to end-users
type SmsClient interface {
	SendSms(c context.Context, destinationNumber string, msgPayload string) error
}

type smsClientFactory func() SmsClient

// New provides an environment specific implementation
var New smsClientFactory
