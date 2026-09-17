package evasion

import (
	"time"
)

// Bridge connects evasion to credential modules
type Bridge struct{}

func NewBridge() *Bridge {
	return &Bridge{}
}

func (b *Bridge) Connect() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "bridge:evasion_to_credential")
	return results, nil
}

func (b *Bridge) Name() string { return "Bridge" }
func (b *Bridge) Timestamp() time.Time { return time.Now() }
