package mailer

import (
	"bytes"
	"text/template"
)

func constructTemplate(templateFile string, data any) (string, string, error) {
	temp, err := template.ParseFS(FS, templateFile)
	if err != nil {
		return "", "", err
	}

	subject := new(bytes.Buffer)
	err = temp.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return "", "", err
	}
	body := new(bytes.Buffer)
	err = temp.ExecuteTemplate(body, "body", data)
	if err != nil {
		return "", "", err
	}
	return subject.String(), body.String(), nil
}
