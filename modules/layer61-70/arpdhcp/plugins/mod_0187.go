package arpdhcp

import (
    "time"
)

type arpdhcp0187 struct{}

func Newarpdhcp0187() *arpdhcp0187 {
    return &arpdhcp0187{}
}

func (e *arpdhcp0187) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0187) Name() string { return "arpdhcp0187" }
func (e *arpdhcp0187) Timestamp() time.Time { return time.Now() }
