package main

import (
	"github.com/baydogan/design-patterns/factorymethod-example2/config"
	"github.com/baydogan/design-patterns/factorymethod-example2/notification"
)

func main() {
	cfg, err := config.LoadConfig("config.json")

	if err != nil {
		panic("Config cannot loaded: " + err.Error())
	}

	sender, err := notification.CreateSender(cfg)
	if err != nil {
		panic("Sender cannot created: " + err.Error())
	}

	sender.Notify("user789", "Factory pattern with config is working!")
}
