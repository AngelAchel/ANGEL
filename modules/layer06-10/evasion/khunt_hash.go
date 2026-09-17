package evasion

import (
	"time"
)

type KhuntHash struct{}

func NewKhuntHash() *KhuntHash {
	return &KhuntHash{}
}

func (e *KhuntHash) Hunt() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "khunt_hash:done")
	return results, nil
}

func (e *KhuntHash) Name() string         { return "KhuntHash" }
func (e *KhuntHash) Timestamp() time.Time { return time.Now() }
