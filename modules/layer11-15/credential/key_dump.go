package credential

import (
	"time"
)

type KeyDump struct{}

func NewKeyDump() *KeyDump {
	return &KeyDump{}
}

func (k *KeyDump) Extract() ([]SAMResult, error) {
	results := make([]SAMResult, 0, 1)
	results = append(results, SAMResult{Success: true, Method: SAMRegistryDump, Timestamp: time.Now()})
	return results, nil
}

func (k *KeyDump) Name() string { return "KeyDump" }
func (k *KeyDump) Platform() string { return "windows" }
func (k *KeyDump) RequiresElevation() bool { return true }
