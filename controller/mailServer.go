package controller

import (
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"log"
	"net/smtp"
)

func SendPasscodeMail(recipientAddr string, passcode string) error {
	mailS := model.NewMailServer()
	to := []string{recipientAddr}

	subject := "Passcode for signing up to What's the Best Keyboard website"
	body := passcode + " is the sign-up passcode for What's the Best Keyboard website\r\nThanks"
	message := []byte(
		"To: " + to[0] + "\r\n" +
			"Subject: " + subject + "\r\n" + "\r\n" +
			body + "\r\n")

	// set up the authentication
	auth := smtp.PlainAuth("", mailS.SmtpFrom, mailS.SmtpPassword, mailS.SmtpHOST)

	// send the email
	err := smtp.SendMail(mailS.SmtpHOST+":"+mailS.SmtpPort, auth, mailS.SmtpFrom, to, message)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	// fmt.Println("Success in sending passcode to a new user")
	return nil
}
