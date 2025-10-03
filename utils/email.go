package utils

import (
	"fmt"
	"net/smtp"
)

func SendContactEmail(
	from, password, to,
	userName, userEmail, userMsg string,
) error {
	// ----- Build the email -------------------------------------------------
	subject := fmt.Sprintf("New message from %s", userName)
	body := fmt.Sprintf(
		"You have a new message from %s (%s):\n\n%s",
		userName, userEmail, userMsg,
	)

	// RFC-5322 style headers (CRLF = \r\n)
	message := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Reply-To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+ // blank line separates headers from body
			"%s",
		from, to, userEmail, subject, body,
	))

	// ----- SMTP configuration ---------------------------------------------
	smtpHost := "smtp.gmail.com"
	smtpPort := "587" // STARTTLS
	addr := smtpHost + ":" + smtpPort

	// PlainAuth works for Gmail when using an App Password
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// ----- Send -----------------------------------------------------------
	return smtp.SendMail(addr, auth, from, []string{to}, message)
}
