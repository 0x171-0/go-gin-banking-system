package service

// NotificationSender defines the interface for different notification strategies
type NotificationSender interface {
	Send(to, code string) error
	GetType() string
}

// NotificationService manages different notification strategies
type NotificationService interface {
	SendVerificationCode(notificationType, to, code string) error
	GetSender(notificationType string) (NotificationSender, error)
	RegisterSender(sender NotificationSender)
	GetAvailableTypes() []string
	// WaitForCompletion waits for all asynchronous notification operations to complete
	WaitForCompletion()
}