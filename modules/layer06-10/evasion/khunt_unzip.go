package evasion

import (
	"time"
)

type KhuntUnzip struct{}

func NewKhuntUnzip() *KhuntUnzip {
	return &KhuntUnzip{}
}

func (e *KhuntUnzip) Hunt() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "khunt_unzip:done")
	return results, nil
}

func (e *KhuntUnzip) Name() string { return "KhuntUnzip" }
func (e *KhuntUnzip) Timestamp() time.Time { return time.Now() }
