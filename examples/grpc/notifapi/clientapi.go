
package notifapi

import (
	"context"
)

type Notifier interface {
	Open(addressPort string) error
	SendEmail(ctx context.Context, recipientEmailAddress, subject, body string) (string,error)
	SendSms(ctx context.Context, recipientPhoneNumber, body string) (string,error)
	GetStatus(ctx context.Context, msgUID string) (NotificationStatus,error)
	Close()
}
