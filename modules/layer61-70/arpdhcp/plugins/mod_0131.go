package arpdhcp

import (
    "time"
)

type arpdhcp0131 struct{}

func Newarpdhcp0131() *arpdhcp0131 {
    return &arpdhcp0131{}
}

func (e *arpdhcp0131) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0131) Name() string { return "arpdhcp0131" }
func (e *arpdhcp0131) Timestamp() time.Time { return time.Now() }
