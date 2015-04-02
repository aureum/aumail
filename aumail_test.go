package aumail

import (
	"os"
	"testing"
)

func TestSend(t *testing.T) {
	var au AuMail

	au.SendGridUser = os.Getenv("SendGridUser")
	au.SendGridKey = os.Getenv("SendGridKey")

	au.From = "abhi@aureum.io"
	emails := make([]string, 0)
	emails = append(emails, "abhi@aureum.io")
	au.Emails = emails

	au.Subject = "Hello"
	au.Text = "Hello, World"
	status, desc := au.Send()

	if !status {
		t.Error("Email did not send " + desc)
	}
}
