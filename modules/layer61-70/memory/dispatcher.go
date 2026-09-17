package memory

import (
	"time"
)

// Dispatcher handles cross-layer dispatch
type Dispatcher struct{}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) Dispatch(layer string, payload []byte) error {
	return nil
}

func (d *Dispatcher) Name() string { return "Dispatcher" }
func (d *Dispatcher) Timestamp() time.Time { return time.Now() }
