package notification

import "fmt"

type NotificationSender interface {
	Notify(userId string, message string)
}

type EmailSender struct {
	SMTPServer string
	From       string
}

func (e EmailSender) Notify(userID string, message string) {
	fmt.Printf("[Email] From: %s, To: %s, SMTP: %s, Message: %s\n",
		e.From, userID, e.SMTPServer, message)
}

type SMSSender struct {
	APIKey string
}

func (s SMSSender) Notify(userID string, message string) {
	fmt.Printf("[SMS] To: %s, Using APIKey: %s, Message: %s\n",
		userID, s.APIKey, message)
}
