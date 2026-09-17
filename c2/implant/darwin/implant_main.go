package implant

import (
	"time"
)

// Darwin implant main entry point
type DarwinMain struct{}

func NewDarwinMain() *DarwinMain {
	return &DarwinMain{}
}

func (m *DarwinMain) Run() error {
	return nil
}

func (m *DarwinMain) Name() string         { return "DarwinMain" }
func (m *DarwinMain) Timestamp() time.Time { return time.Now() }
