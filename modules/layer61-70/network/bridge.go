package network

import (
	"time"
)

// Bridge connects layer 61-70 back to C2
type Bridge struct{}

func NewBridge() *Bridge {
	return &Bridge{}
}

func (b *Bridge) Connect() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "bridge:v31_to_c2")
	return results, nil
}

func (b *Bridge) Name() string { return "Bridge" }
func (b *Bridge) Timestamp() time.Time { return time.Now() }
