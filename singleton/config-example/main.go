package main

import (
	"fmt"
	"github.com/baydogan/design-patterns/singleton-example3/config"
	"log"
)

func main() {
	cfg := config.GetInstance()

	log.Println("Application started")
	log.Printf("Loaded config: AppName=%s, Version=%s\n", cfg.AppName, cfg.Version)

	cfg2 := config.GetInstance()

	fmt.Printf("cfg pointer address: %p\n", cfg)
	fmt.Printf("cfg2 pointer address: %p\n", cfg2)

	fmt.Println("cfg == cfg2 ?", cfg == cfg2)
}
