// emailapi provides an easy interface to send out emails
package emailapi

// Send is used to send out a plain text email.
// It can be overridden because it some environments or situations it is not appropriate
var Send func(recipient, subject, body string) = defaultSend

// SetDefaultSend is useful to switching back to the default implementation (after a test has completed)
func SetDefaultSend() {
	Send = defaultSend
}

// defaultSend is used to send out a plain text email
func defaultSend(recipient, subject, body string) {
	// TODO send email using some external service
}
