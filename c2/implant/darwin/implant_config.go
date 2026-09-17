package implant

import (
	"encoding/json"
	"time"
)

// Darwin implant configuration
type DarwinConfig struct {
	ImplantID   string        `json:"implant_id"`
	C2URL       string        `json:"c2_url"`
	Sleep       time.Duration `json:"sleep"`
	Jitter      int           `json:"jitter"`
	Retry       int           `json:"retry"`
	Proxy       bool          `json:"proxy"`
	Evasion     bool          `json:"evasion"`
	Persistence bool          `json:"persistence"`
}

func NewDarwinConfig() *DarwinConfig {
	return &DarwinConfig{
		ImplantID:   "darwin-001",
		C2URL:       "",
		Sleep:       60,
		Jitter:      10,
		Retry:       3,
		Proxy:       false,
		Evasion:     true,
		Persistence: true,
	}
}

func (c *DarwinConfig) Validate() error          { return nil }
func (c *DarwinConfig) Marshal() ([]byte, error) { return json.Marshal(c) }
func (c *DarwinConfig) Name() string             { return "DarwinConfig" }
