package emailsending

import (
	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/api/emailsender"
)

type noopEmailSender struct{}

func New() emailsender.EmailSender {
	return &noopEmailSender{}
}

func (ss *noopEmailSender) SendEmail(address string, subject, body string) error {
	return nil
}
