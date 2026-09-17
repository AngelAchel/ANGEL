package implant

import (
	"encoding/json"
	"time"
)

// Windows implant configuration
type WindowsConfig struct {
	ImplantID   string        `json:"implant_id"`
	C2URL       string        `json:"c2_url"`
	Sleep       time.Duration `json:"sleep"`
	Jitter      int           `json:"jitter"`
	Retry       int           `json:"retry"`
	Proxy       bool          `json:"proxy"`
	Evasion     bool          `json:"evasion"`
	Persistence bool          `json:"persistence"`
}

func NewWindowsConfig() *WindowsConfig {
	return &WindowsConfig{
		ImplantID:   "windows-001",
		C2URL:       "",
		Sleep:       60,
		Jitter:      10,
		Retry:       3,
		Proxy:       false,
		Evasion:     true,
		Persistence: true,
	}
}

func (c *WindowsConfig) Validate() error          { return nil }
func (c *WindowsConfig) Marshal() ([]byte, error) { return json.Marshal(c) }
func (c *WindowsConfig) Name() string             { return "WindowsConfig" }
