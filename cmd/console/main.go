package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/angel-platform/angel/gateway"
)

func main() {
	addr := flag.String("addr", "0.0.0.0", "Bind address")
	port := flag.Int("port", 3000, "Listen port")
	jwtSecret := flag.String("jwt-secret", "", "JWT signing secret")
	flag.Parse()

	if *jwtSecret == "" {
		*jwtSecret = os.Getenv("JWT_SECRET")
	}
	if *jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable must be set")
	}

	cfg := &gateway.Config{
		Addr:        *addr,
		Port:        *port,
		JWTSecret:   *jwtSecret,
		EnableRBAC:  true,
		EnableMetrics: true,
	}

	gw := gateway.New(cfg)
	fmt.Printf("API Gateway starting on %s:%d\n", *addr, *port)
	if err := gw.Start(); err != nil {
		log.Fatalf("Gateway failed: %v", err)
	}
}