package service

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"

	"go-gin-template/api/config"
)

// EmailSender implements NotificationSender for email notifications
type EmailSender struct{
	// wg is used to wait for all email sending goroutines to complete
	wg sync.WaitGroup
}

func NewEmailSender() NotificationSender {
	return &EmailSender{}
}

func (e *EmailSender) Send(to, code string) error {
	// Check email configuration before starting the goroutine
	config := config.GetEmailConfig()
	if config.SMTPUsername == "" || config.SMTPPassword == "" {
		return fmt.Errorf("email credentials not configured")
	}
	
	// Increment the wait group counter
	e.wg.Add(1)
	
	// Launch a goroutine to send the email
	go func(recipient, verificationCode string) {
		// Ensure the wait group counter is decremented when the goroutine completes
		defer e.wg.Done()
		
		subject := "Transfer Verification Code"
		body := fmt.Sprintf("Your verification code is: %s\nThis code will expire in 5 minutes.", verificationCode)

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
			[]string{recipient},
			[]byte(emailBody),
		)

		if err != nil {
			log.Printf("Failed to send email: %v", err)
			return
		}

		log.Printf("📧 Email sent successfully to %s", recipient)
	}(to, code)
	
	// Return immediately, not waiting for the email to be sent
	// The caller can continue execution while the email is being sent in the background
	log.Printf("📧 Email sending initiated to %s", to)
	return nil
}

func (e *EmailSender) GetType() string {
	return "email"
}

// WaitForCompletion waits for all email sending goroutines to complete
// This can be called when the application is shutting down to ensure all emails are sent
func (e *EmailSender) WaitForCompletion() {
	e.wg.Wait()
	log.Printf("All email sending operations completed")
}