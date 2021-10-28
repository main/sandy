package mailer

import (
	"bytes"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"html/template"
	"log"
)

type Options struct {
	InstanceName string

	TemplateEmailSilenceMaxTimeExceeded   string
	TemplateEmailOperationMaxTimeExceeded string

	SenderEmail string
	SenderName  string
	Receivers   []string

	SendGridKey string
}

type Mailer struct {
	options Options
}

func NewMailer(options Options) *Mailer {
	return &Mailer{
		options: options,
	}
}

func (m Mailer) SendMaxSilenceEmails(templateArgs map[string]interface{}) error {
	for _, recipientEmail := range m.options.Receivers {
		if err := m.sendmail(recipientEmail,
			fmt.Sprintf("Sandy: Instance %s has reached max silence time", m.options.InstanceName),
			render(m.options.TemplateEmailSilenceMaxTimeExceeded,
				templateArgs),
		); err != nil {
			return err
		}
	}

	return nil
}

func (m Mailer) SendMaxOperationEmails(templateArgs map[string]interface{}) error {
	for _, recipientEmail := range m.options.Receivers {
		if err := m.sendmail(recipientEmail,
			fmt.Sprintf("Sandy: Instance %s has exceeded max operation time", m.options.InstanceName),
			render(m.options.TemplateEmailOperationMaxTimeExceeded,
				templateArgs),
		); err != nil {
			return err
		}
	}

	return nil
}

func (m Mailer) from() *mail.Email {
	return mail.NewEmail(m.options.SenderName, m.options.SenderEmail)
}

func recipient(email string) *mail.Email {
	return mail.NewEmail("", email)
}

func (m Mailer) sendmail(recipientEmail, subject, body string) error {
	msg := mail.NewSingleEmail(m.from(), subject, recipient(recipientEmail), body, fmt.Sprintf("<pre>%s</pre>", body))
	cli := sendgrid.NewSendClient(m.options.SendGridKey)
	resp, err := cli.Send(msg)

	if err != nil {
		log.Println("send email status:", resp.StatusCode)
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("error sending message. status %d text:%s", resp.StatusCode, resp.Body) // TODO: error handling
	}

	return nil
}

func render(text string, data interface{}) string {
	tpl, err := template.New("mail").Parse(text)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer([]byte{})
	if err := tpl.Execute(buf, data); err != nil {
		panic(err)
	}
	return buf.String()
}
