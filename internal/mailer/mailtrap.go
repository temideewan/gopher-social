package mailer

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/wneessen/go-mail"
)

type MailtrapMailer struct {
	fromEmail string
	apiKey    string
	client    *mail.Client
}

func New(host string, port int, username, password, sender string) (*MailtrapMailer, error) {

	client, err := mail.NewClient(
		host,
		mail.WithSMTPAuth(mail.SMTPAuthLogin),
		mail.WithPort(port),
		mail.WithUsername(username),
		mail.WithPassword(password),
		mail.WithTimeout(5*time.Second),
	)

	if err != nil {
		return nil, err
	}

	mailer := &MailtrapMailer{
		client:    client,
		fromEmail: sender,
	}
	return mailer, nil
}

func (m *MailtrapMailer) Send(templateFile, username, email string, data any, isSandbox bool) error {
	msg := mail.NewMsg()

	err := msg.To(email)
	if err != nil {
		return err
	}

	err = msg.From(FromName)
	if err != nil {
		return nil
	}
	// template parsing
	subject := new(bytes.Buffer)
	plainBody := new(bytes.Buffer)
	htmlBody := new(bytes.Buffer)

	msg.Subject(subject.String())
	msg.SetBodyString(mail.TypeTextPlain, plainBody.String())
	msg.AddAlternativeString(mail.TypeTextHTML, htmlBody.String())

	for i := 0; i < maxRetries; i++ {
		err = m.client.DialAndSend(msg)
		if err != nil {
			log.Printf("Failed to send email to %v, attempt %d of %d", email, i+1, maxRetries)
			log.Printf("Error :%v", err.Error())

			// exponential backoff
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		log.Printf("Email send successfully to %v", email)
		return nil
	}
	return fmt.Errorf("Email sending failed")
}
