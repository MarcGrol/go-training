package emailsender

//go:generate mockgen -source=emailSender.go -destination=emailSenderMocks.go -package=emailsender EmailSender

type EmailSender interface {
	SendEmail(emailAddress, subject, body string) error
}
