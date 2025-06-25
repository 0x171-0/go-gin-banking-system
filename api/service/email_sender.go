package service

import (
	"fmt"
	"log"
)

// EmailSender implements NotificationSender for email notifications
type EmailSender struct{}

func NewEmailSender() NotificationSender {
	return &EmailSender{}
}

func (e *EmailSender) Send(to, code string) error {
	subject := "Transfer Verification Code"
	body := fmt.Sprintf("Your verification code is: %s\nThis code will expire in 5 minutes.", code)
	
	// TODO: Implement actual email sending using providers like:
	// - AWS SES
	// - SendGrid
	// - SMTP
	
	log.Printf("Sending email to %s: Subject: %s, Body: %s", to, subject, body)
	fmt.Printf("📧 Email sent to %s\nSubject: %s\nBody: %s\n", to, subject, body)
	
	return nil
}

func (e *EmailSender) GetType() string {
	return "email"
}