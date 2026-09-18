package sqlinject

import (
	"time"

	"github.com/angel-platform/angel/modules/eventbus"
)

// Dispatcher handles cross-layer dispatch
type Dispatcher struct{}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) Dispatch(layer string, payload []byte) error {
	eventbus.Publish(eventbus.NewEvent("dispatch", "sqlinject", "*", "event", nil))
	eventbus.Publish(eventbus.NewEvent("dispatch", "sqlinject", "*", "event", nil))
	return nil
}

func (d *Dispatcher) Name() string         { return "Dispatcher" }
func (d *Dispatcher) Timestamp() time.Time { return time.Now() }
