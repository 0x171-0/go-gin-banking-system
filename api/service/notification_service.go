package service

import (
	"fmt"
)

// notificationService implements NotificationService
type notificationService struct {
	senders map[string]NotificationSender
}

func NewNotificationService() NotificationService {
	service := &notificationService{
		senders: make(map[string]NotificationSender),
	}
	
	// Register email sender
	emailSender := NewEmailSender()
	service.senders[emailSender.GetType()] = emailSender
	
	// Register SMS sender
	smsSender := NewSMSSender()
	service.senders[smsSender.GetType()] = smsSender
	
	return service
}

func (n *notificationService) SendVerificationCode(notificationType, to, code string) error {
	sender, err := n.GetSender(notificationType)
	if err != nil {
		return err
	}
	
	return sender.Send(to, code)
}

func (n *notificationService) GetSender(notificationType string) (NotificationSender, error) {
	sender, exists := n.senders[notificationType]
	if !exists {
		return nil, fmt.Errorf("notification type '%s' is not supported", notificationType)
	}
	
	return sender, nil
}

func (n *notificationService) RegisterSender(sender NotificationSender) {
	n.senders[sender.GetType()] = sender
}

func (n *notificationService) GetAvailableTypes() []string {
	types := make([]string, 0, len(n.senders))
	for senderType := range n.senders {
		types = append(types, senderType)
	}
	return types
}

// WaitForCompletion waits for all asynchronous notification operations to complete
func (n *notificationService) WaitForCompletion() {
	// Wait for email sender to complete all operations
	if emailSender, ok := n.senders["email"].(*EmailSender); ok {
		emailSender.WaitForCompletion()
	}
	
	// If there are other asynchronous senders, wait for them here
	// For example, if SMS sender also uses goroutines:
	// if smsSender, ok := n.senders["sms"].(*SMSSender); ok {
	//     smsSender.WaitForCompletion()
	// }
}