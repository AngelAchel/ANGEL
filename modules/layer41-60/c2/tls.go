package c2

import (
	"time"
)

type TLS struct{}

func NewTLS() *TLS {
	return &TLS{}
}

func (e *TLS) Handshake(target string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "tls:handshaked")
	return results, nil
}

func (e *TLS) Name() string { return "TLS" }
func (e *TLS) Category() C2Category { return CategoryC2 }
func (e *TLS) Timestamp() time.Time { return time.Now() }
