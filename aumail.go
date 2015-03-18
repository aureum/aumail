package aumail

import (
	"net/mail"

	"github.com/sendgrid/sendgrid-go"
)

type AuMail struct {
	SendGridUser string
	SendGridKey  string
	From         string
	Emails       []string
	Subject      string
	Text         string
	HTML         string
}

func (au *AuMail) Send() (ifError bool, desc string) {
	sg := sendgrid.NewSendGridClient(au.SendGridUser, au.SendGridKey)
	message := sendgrid.NewMail()
	for i := 0; i < len(au.Emails); i++ {
		address, _ := mail.ParseAddress(au.Emails[i])
		message.AddRecipient(address)
	}

	message.SetSubject(au.Subject)
	if au.Text != "" {
		message.SetText(au.Text)
	} else {
		message.SetHTML(au.HTML)
	}

	message.SetFrom(au.From)
	if r := sg.Send(message); r == nil {
		return true, ""
	} else {
		return false, r.Error()
	}
}
