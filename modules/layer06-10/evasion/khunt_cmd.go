package evasion

import (
	"time"
)

type KhuntCmd struct{}

func NewKhuntCmd() *KhuntCmd {
	return &KhuntCmd{}
}

func (e *KhuntCmd) Hunt() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "khunt_cmd:done")
	return results, nil
}

func (e *KhuntCmd) Name() string { return "KhuntCmd" }
func (e *KhuntCmd) Timestamp() time.Time { return time.Now() }
