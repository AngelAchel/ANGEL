package crypto

import (
	"time"
)

// Bridge connects additional modules v1 to v2
type Bridge struct{}

func NewBridge() *Bridge {
	return &Bridge{}
}

func (b *Bridge) Connect() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "bridge:v1_to_v2")
	return results, nil
}

func (b *Bridge) Name() string { return "Bridge" }
func (b *Bridge) Timestamp() time.Time { return time.Now() }
