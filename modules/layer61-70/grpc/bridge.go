package grpc

import (
	"time"

	"github.com/angel-platform/angel/modules/eventbus"
)

// Bridge connects layers
type Bridge struct{}

func NewBridge() *Bridge {
	return &Bridge{}
}

func (b *Bridge) Connect() error {
	eventbus.Subscribe("layer:*")
	return nil
}

func (b *Bridge) Name() string         { return "Bridge" }
func (b *Bridge) Timestamp() time.Time { return time.Now() }
