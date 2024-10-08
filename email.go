package main

import (
	"fmt"
	ht "html/template"
	"log"
	_ "text/template"

	"github.com/wneessen/go-mail"
)

func SendEmail(cfg Config, message *mail.Msg) {
	c, err := mail.NewClient(cfg.EMAIL_HOST, mail.WithPort(cfg.EMAIL_PORT), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.EMAIL_USER), mail.WithPassword(cfg.EMAIL_PASSWORD))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}
	defer c.Close()

	if err := c.DialAndSend(message); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}

	fmt.Println("Email Sent!") // Send email here
}

func CreateMail(cfg Config) (*mail.Msg, error) {
	m := mail.NewMsg()
	if err := m.From("dave.sender@example.com"); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To("dave.receiver@example.com"); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}

	// textBodyTemplate, err := ReadFile("email/test/template.txt")
	// if err != nil {
	// 	return nil, err
	// }

	// ttpl, err := tt.New("texttpl").Parse(textBodyTemplate)
	// if err != nil {
	// 	return nil, err
	// }

	htmlBodyTemplate, err := ReadFile("email/test/template.html")
	if err != nil {
		return nil, err
	}

	htpl, err := ht.New("htmltpl").Parse(htmlBodyTemplate)
	if err != nil {
		return nil, err
	}

	m.Subject("Hello, Gophers!")
	err = m.SetBodyHTMLTemplate(htpl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to set body: %s", err)
	}

	return m, nil
}

func EmailOnTimer(cfg Config) {
	m, err := CreateMail(cfg)
	if err != nil {
		log.Fatalf("failed to create mail: %s", err)
	}

	SendEmail(cfg, m)
}
