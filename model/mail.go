package model

import (
	"github.com/joho/godotenv"
	"net/smtp"
	"os"
)

type Mail struct {
	smtpHOST     string
	smtpPort     string
	smtpFrom     string
	smtpPassword string
}

func NewMailServer() Mail {
	godotenv.Load()
	return Mail{
		smtpHOST:     os.Getenv("SMTPHOST"),
		smtpPort:     os.Getenv("SMTPPORT"),
		smtpFrom:     os.Getenv("SMTPFROM"),
		smtpPassword: os.Getenv("SMTPPASS"),
	}
}

// SendMail sends a given message to a given recipient
func (mail Mail) SendMail(to []string, message []byte) error {
	// set up the authentication
	auth := smtp.PlainAuth("", mail.smtpFrom, mail.smtpPassword, mail.smtpHOST)
	// send the email
	err := smtp.SendMail(mail.smtpHOST+":"+mail.smtpPort, auth, mail.smtpFrom, to, message)
	if err != nil {
		return err
	}
	return nil
}
