package service

import (
	"fmt"
	"log"
	"net/smtp"

	"go-gin-template/api/config"
)

// EmailSender implements NotificationSender for email notifications
type EmailSender struct{}

func NewEmailSender() NotificationSender {
	return &EmailSender{}
}

func (e *EmailSender) Send(to, code string) error {
	subject := "Transfer Verification Code"
	body := fmt.Sprintf("Your verification code is: %s\nThis code will expire in 5 minutes.", code)

	config := config.GetEmailConfig()

	if config.SMTPUsername == "" || config.SMTPPassword == "" {
		return fmt.Errorf("email credentials not configured")
	}

	// Set up authentication information
	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPHost)

	// Compose email
	mime := "MIME-version: 1.0;\n" +
		"Content-Type: text/plain; charset=\"UTF-8\";\n\n"
	emailBody := fmt.Sprintf("Subject: %s\n%s\n%s", subject, mime, body)

	// Send email
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", config.SMTPHost, config.SMTPPort),
		auth,
		config.SMTPUsername,
		[]string{to},
		[]byte(emailBody),
	)

	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Printf("📧 Email sent successfully to %s", to)
	return nil
}

func (e *EmailSender) GetType() string {
	return "email"
}