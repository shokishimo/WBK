package controller

import (
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"log"
	"net/smtp"
)

func SendPasscodeMail(recipientAddr string) error {
	mailS := model.NewMailServer()
	to := []string{recipientAddr}

	message := []byte("BestKeyboard website message.")

	// set up the authentication
	auth := smtp.PlainAuth("", mailS.SmtpFrom, mailS.SmtpPassword, mailS.SmtpHOST)

	// send the email
	err := smtp.SendMail(mailS.SmtpHOST+":"+mailS.SmtpPort, auth, mailS.SmtpFrom, to, message)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}
