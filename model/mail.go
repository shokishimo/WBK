package model

import (
	"github.com/joho/godotenv"
	"os"
)

type Mail struct {
	SmtpHOST     string
	SmtpPort     string
	SmtpFrom     string
	SmtpPassword string
}

func NewMailServer() Mail {
	godotenv.Load()
	return Mail{
		SmtpHOST:     os.Getenv("SMTPHOST"),
		SmtpPort:     os.Getenv("SMTPPORT"),
		SmtpFrom:     os.Getenv("SMTPFROM"),
		SmtpPassword: os.Getenv("SMTPPASS"),
	}
}
