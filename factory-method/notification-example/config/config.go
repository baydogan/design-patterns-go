package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	NotificationType string        `json:"notification_type"`
	EmailConfig      EmailSettings `json:"email"`
	SMSConfig        SMSSettings   `json:"sms"`
}

type EmailSettings struct {
	SMTPServer string `json:"smtp_server"`
	From       string `json:"from"`
}

type SMSSettings struct {
	APIKey string `json:"api_key"`
}

func LoadConfig(path string) (*AppConfig, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var cfg AppConfig
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
