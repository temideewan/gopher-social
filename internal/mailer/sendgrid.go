package mailer

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridMailer struct {
	fromEmail string
	apiKey string
	client *sendgrid.Client
}

func NewSendGrid(apiKey, fromEmail string) *SendGridMailer {
	client := sendgrid.NewSendClient(apiKey)

	return &
}
func (m *SenSendGridMailer) Send(templateFile, username, email string, data any, isSandbox bool) error {
	from := mail.NewEmail
}
