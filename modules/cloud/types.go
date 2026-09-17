package cloud

import (
	"fmt"
	"time"
)

// EncodeForChannel encodes data for covert channel transmission.
func EncodeForChannel(data string, channel string) string {
	return fmt.Sprintf("%s:%s:%x", channel, data, len(data))
}

// CloudConfig holds configuration for cloud service engines.
type CloudConfig struct {
	Provider     string
	AccessKey    string
	SecretKey    string
	SessionToken string
	Region       string
	ProjectID    string
	TenantID     string
	ClientID     string
	ClientSecret string
	Profile      string
	Endpoint     string
}

// CloudResult holds the result of a cloud covert operation.
type CloudResult struct {
	ID        string    `json:"id"`
	Provider  string    `json:"provider"`
	AccountID string    `json:"account_id"`
	Channel   string    `json:"channel"`
	Payload   string    `json:"payload"`
	Encoded   string    `json:"encoded"`
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
	RiskScore float64   `json:"risk_score"`
	Details   string    `json:"details"`
}
