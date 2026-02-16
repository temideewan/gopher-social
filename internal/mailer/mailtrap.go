package mailer

import (
	"errors"

	gomail "gopkg.in/mail.v2"
)

type mailtrapClient struct {
	fromEmail string
	apiKey    string
	username  string
	password  string
}

func NewMailTrapClient(apiKey, fromEmail, username, password string) (mailtrapClient, error) {
	if apiKey == "" {
		return mailtrapClient{}, errors.New("api key is required")
	}
	return mailtrapClient{
		fromEmail: fromEmail,
		apiKey:    apiKey,
		username:  username,
		password:  password,
	}, nil
}

func (m mailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	// template parsing and building
	subject, body, err := constructTemplate("templates/"+templateFile, data)
	if err != nil {
		return -1, err
	}

	message := gomail.NewMessage()
	message.SetHeader("From", m.fromEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", subject)

	message.AddAlternative("text/html", body)

	dialer := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, m.username, m.password)

	if err := dialer.DialAndSend(message); err != nil {
		return -1, err
	}
	return sendMail(func() (int, error) {
		err := dialer.DialAndSend(message)
		if err != nil {
			return -1, err
		}
		return 200, nil
	})
}
