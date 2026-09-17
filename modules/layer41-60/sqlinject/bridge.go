package sqlinject

import (
	"time"
)

// Bridge connects additional modules v3 to v3.1
type Bridge struct{}

func NewBridge() *Bridge {
	return &Bridge{}
}

func (b *Bridge) Connect() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "bridge:v3_to_v31")
	return results, nil
}

func (b *Bridge) Name() string         { return "Bridge" }
func (b *Bridge) Timestamp() time.Time { return time.Now() }
