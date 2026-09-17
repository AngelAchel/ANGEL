package webmisc

import (
	"time"
)

// Bridge connects layers
type Bridge struct{}

func NewBridge() *Bridge {
	return &Bridge{}
}

func (b *Bridge) Connect() error {
	return nil
}

func (b *Bridge) Name() string         { return "Bridge" }
func (b *Bridge) Timestamp() time.Time { return time.Now() }
