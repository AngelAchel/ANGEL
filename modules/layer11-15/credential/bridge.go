package credential

import (
	"time"
)

// Bridge connects credential to infrastructure modules
type Bridge struct{}

func NewBridge() *Bridge {
	return &Bridge{}
}

func (b *Bridge) Connect() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "bridge:credential_to_infra")
	return results, nil
}

func (b *Bridge) Name() string { return "Bridge" }
func (b *Bridge) Timestamp() time.Time { return time.Now() }
