package controller

import (
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
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

	err := mailS.SendMail(to, message)
	if err != nil {
		return err
	}
	return nil
}
