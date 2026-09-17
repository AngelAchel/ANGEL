package implant

import (
	"time"
)

// Windows implant registration
type WindowsRegister struct{}

func NewWindowsRegister() *WindowsRegister {
	return &WindowsRegister{}
}

func (r *WindowsRegister) Register() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "windows:registered")
	return results, nil
}

func (r *WindowsRegister) Name() string { return "WindowsRegister" }
func (r *WindowsRegister) Timestamp() time.Time { return time.Now() }
