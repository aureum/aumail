package aumail

import "testing"

func TestSend(t *testing.T) {
	var au AuMail

	au.SendGridUser = ""
	au.SendGridKey = ""

	au.From = "abhi@aureum.io"
	emails := make([]string, 0)
	emails = append(emails, "cheryl@aureum.io")
	au.Emails = emails

	au.Subject = "Hello"
	au.Text = "Hello, World"
	status, desc := au.Send()

	if !status {
		t.Error("Email did not send " + desc)
	}
}
