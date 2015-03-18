package aumail

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	var au aumail

	au.SendGridUser = ""
	au.SendGridKey = ""

	emails := make([]string, 0)
	emails = append(emails, "cheryl@aureum.io")
	au.Emails = emails
	au.From = "abhi@aureum.io"

	au.Subject = "Hello"
	au.Text = "Hello, World"
	fmt.Println(au.Send())
}
