package controller

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shokishimo/WhatsTheBestKeyboard/model"
	"os"
)

func SendPasscodeMail(recipientAddr string, passcode string) error {
	mail := model.NewMailServer()
	to := []string{recipientAddr}

	subject := "Passcode for signing up to What's the Best Keyboard website"
	body := passcode + " is the sign-up passcode for What's the Best Keyboard website\r\nThanks"
	message := []byte(
		"To: " + to[0] + "\r\n" +
			"Subject: " + subject + "\r\n" + "\r\n" +
			body + "\r\n")

	err := mail.SendMail(to, message)
	if err != nil {
		return err
	}
	return nil
}

func SendKBRequestMailToOwner(nickname string, keyboard string, url string) error {
	mail := model.NewMailServer()
	to := os.Getenv("OWNER_EMAIL")
	if to == "" {
		return errors.New("Error: os.Getenv(\"OWNER_EMAIL\")")
	}
	subject := fmt.Sprintf("WBK: New Keyboard Request from %s", nickname)
	body := fmt.Sprintf("From %s \r\n Keyboard Name: %s \r\n Keyboard URL: %s \r\n\r\n Thanks", nickname, keyboard, url)
	message := []byte(
		"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" + "\r\n" +
			body + "\r\n")

	if err := mail.SendMail([]string{to}, message); err != nil {
		return err
	}
	return nil
}
