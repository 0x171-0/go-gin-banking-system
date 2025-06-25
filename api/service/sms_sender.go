package service

import (
	"fmt"
	"log"
)

// SMSSender implements NotificationSender for SMS notifications
type SMSSender struct{}

func NewSMSSender() NotificationSender {
	return &SMSSender{}
}

func (s *SMSSender) Send(to, code string) error {
	message := fmt.Sprintf("Your verification code is: %s. This code will expire in 5 minutes.", code)
	
	// TODO: Implement actual SMS sending using providers like:
	// - Twilio
	// - AWS SNS
	// - Vonage (Nexmo)
	// - Firebase Cloud Messaging
	
	log.Printf("Sending SMS to %s: %s", to, message)
	fmt.Printf("📱 SMS sent to %s\nMessage: %s\n", to, message)
	
	return nil
}

func (s *SMSSender) GetType() string {
	return "sms"
}