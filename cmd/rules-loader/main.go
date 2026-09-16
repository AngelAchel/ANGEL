package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/angel-platform/angel/pkg/supabase"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	engine := supabase.NewRulesEngine()

	if err := engine.LoadAllRules(); err != nil {
		log.Printf("Warning: Supabase rules unavailable, loading local fallback: %v", err)
		engine.LoadLocalRules("/app/data/rules.json")
	}

	categories := []string{"c2", "exploit", "auth", "network"}
	for _, cat := range categories {
		rules := engine.GetRules(cat)
		fmt.Printf("Category: %s - %d rules\n", cat, len(rules))
		for _, rule := range rules {
			fmt.Printf("  [%d] %s: %s\n", rule.Priority, rule.Name, rule.Pattern)
		}
	}

	fmt.Println("\nRules loaded successfully!")

	// Run as daemon — reload rules every 60s
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Rules loader running as daemon. Press Ctrl+C to stop.")
	for {
		select {
		case <-ticker.C:
			log.Println("Reloading rules...")
			if err := engine.LoadAllRules(); err != nil {
				log.Printf("Warning: Supabase rules unavailable, loading local fallback: %v", err)
				engine.LoadLocalRules("/app/data/rules.json")
			}
			log.Println("Rules reloaded successfully.")
		case <-sigChan:
			log.Println("Shutting down rules loader...")
			os.Exit(0)
		}
	}

}
