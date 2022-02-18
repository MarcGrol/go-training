package smssender

//go:generate mockgen -source=smsSender.go -destination=smsSenderMocks.go -package=smssender SmsSender

type SmsSender interface {
	SendSms(phoneNumber string, smsContent string) error
}
