package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/angel-platform/angel/modules/layer01-05/c2server"
)

func main() {
	bind := flag.String("bind", "0.0.0.0", "Bind address")
	port := flag.Int("port", 8443, "Listen port")
	cryptoKey := flag.String("key", "", "Encryption key")
	maxAgents := flag.Int("max-agents", 100, "Max agents")
	flag.Parse()

	if *cryptoKey == "" {
		*cryptoKey = os.Getenv("TEAMSERVER_KEY")
	}
	if *cryptoKey == "" {
		log.Fatal("TEAMSERVER_KEY environment variable must be set")
	}

	cfg := &c2server.TeamserverConfig{
		BindAddr:  *bind,
		BindPort:  *port,
		CryptoKey: *cryptoKey,
		MaxAgents: *maxAgents,
		DBPath:    "teamserver.db",
	}

	ts, err := c2server.NewTeamserver(cfg)
	if err != nil {
		log.Fatalf("Failed to create teamserver: %v", err)
	}

	fmt.Printf("Teamserver starting on %s:%d\n", *bind, *port)
	if err := ts.Start(); err != nil {
		log.Fatalf("Teamserver failed: %v", err)
	}

	// Block forever
	select {}
}
