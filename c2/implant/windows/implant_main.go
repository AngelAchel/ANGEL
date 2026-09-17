package implant

import (
	"time"
)

// Windows implant main entry point
type WindowsMain struct{}

func NewWindowsMain() *WindowsMain {
	return &WindowsMain{}
}

func (m *WindowsMain) Run() error {
	return nil
}

func (m *WindowsMain) Name() string { return "WindowsMain" }
func (m *WindowsMain) Timestamp() time.Time { return time.Now() }
