package notification

import (
	"fmt"
	"github.com/baydogan/design-patterns/factorymethod-example2/config"
)

func CreateSender(cfg *config.AppConfig) (NotificationSender, error) {
	switch cfg.NotificationType {
	case "email":
		return EmailSender{
			SMTPServer: cfg.EmailConfig.SMTPServer,
			From:       cfg.EmailConfig.From,
		}, nil
	case "sms":
		return SMSSender{
			APIKey: cfg.SMSConfig.APIKey,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported notification type: %s", cfg.NotificationType)
	}
}
