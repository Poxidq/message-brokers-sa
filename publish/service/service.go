package service

import (
	"log"
	"publish/config"

	gomail "gopkg.in/mail.v2"
)

func SendEmail(subject, body string) error {
	emailRecipients := config.CFG.EmailRecipients
	fromEmail := config.CFG.EmailAddress
	emailPassword := config.CFG.EmailPassword

	m := gomail.NewMessage()

	m.SetHeader("From", fromEmail)
	m.SetHeader("To", emailRecipients)
	m.SetHeader("Subject", subject)

	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, fromEmail, emailPassword)

	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		panic(err)
	}
	return nil
}
