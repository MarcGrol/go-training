package emailsending

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MarcGrol/go-training/examples/registrationServiceGrpc/lib/api/emailsender"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendgridEmailSender struct {
	apiKey string
}

func New() emailsender.EmailSender {

	return &sendgridEmailSender{
		apiKey: os.Getenv("SENDGRID_API_KEY"),
	}
}

func (es *sendgridEmailSender) SendEmail(recipientAddress string, subject, body string) error {
	from := mail.NewEmail("Marc Grol", "mgrol@xebia.com")
	to := mail.NewEmail("", recipientAddress)
	htmlContent := "<strong>" + body + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, body, htmlContent)
	client := sendgrid.NewSendClient(es.apiKey)
	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("Error sending email: %s", err)
	}
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error sending email: %d", response.StatusCode)
	}
	log.Printf("Successfully send mail to %s (%+v	)", recipientAddress, response)

	return nil
}
