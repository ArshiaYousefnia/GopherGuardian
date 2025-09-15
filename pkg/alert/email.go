package alert

import (
	"gopehrguardian/pkg/config"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

const (
	port = 587
	host = "smtp.gmail.com"
)

func emailAlert(target *config.Target, text string) {
	sender := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	receiver := target.Alert.Email

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", "GopherGuardian alert")
	m.SetBody("text/plain", text)

	d := gomail.NewDialer(host, port, sender, password)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err, target.Name, target.Alert.Email)
		return
	}
}
