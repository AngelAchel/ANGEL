package c2

import (
	"time"
)

type TLS struct{}

func NewTLS() *TLS {
	return &TLS{}
}

func (t *TLS) Handshake() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "tls:done")
	return results, nil
}

func (t *TLS) Name() string         { return "TLS" }
func (t *TLS) Timestamp() time.Time { return time.Now() }
