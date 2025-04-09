package config

import (
	"fmt"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

func GetInstance() *Config {
	once.Do(func() {
		fmt.Println("Config instance is being created...")
		instance = &Config{
			AppName: "MyApp",
			Version: "1.0.0",
		}
	})
	return instance
}
