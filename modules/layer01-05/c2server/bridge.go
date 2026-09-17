package c2server

import (
	"time"
)

// Bridge connects C2 core to evasion modules
type Bridge struct{}

func NewBridge() *Bridge {
	return &Bridge{}
}

func (b *Bridge) Connect() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "bridge:c2_to_evasion")
	return results, nil
}

func (b *Bridge) Name() string         { return "Bridge" }
func (b *Bridge) Timestamp() time.Time { return time.Now() }
