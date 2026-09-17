package evasion

import (
	"time"
)

type KhuntFS struct{}

func NewKhuntFS() *KhuntFS {
	return &KhuntFS{}
}

func (e *KhuntFS) Hunt() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "khunt_fs:done")
	return results, nil
}

func (e *KhuntFS) Name() string { return "KhuntFS" }
func (e *KhuntFS) Timestamp() time.Time { return time.Now() }
