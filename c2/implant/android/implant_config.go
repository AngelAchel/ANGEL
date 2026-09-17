package implant

import (
	"encoding/json"
	"time"
)

// Android implant configuration
type AndroidConfig struct {
	ImplantID   string        `json:"implant_id"`
	C2URL       string        `json:"c2_url"`
	Sleep       time.Duration `json:"sleep"`
	Jitter      int           `json:"jitter"`
	Retry       int           `json:"retry"`
	Proxy       bool          `json:"proxy"`
	Evasion     bool          `json:"evasion"`
	Persistence bool          `json:"persistence"`
}

func NewAndroidConfig() *AndroidConfig {
	return &AndroidConfig{
		ImplantID:   "android-001",
		C2URL:       "",
		Sleep:       60,
		Jitter:      10,
		Retry:       3,
		Proxy:       false,
		Evasion:     true,
		Persistence: true,
	}
}

func (c *AndroidConfig) Validate() error          { return nil }
func (c *AndroidConfig) Marshal() ([]byte, error) { return json.Marshal(c) }
func (c *AndroidConfig) Name() string             { return "AndroidConfig" }
