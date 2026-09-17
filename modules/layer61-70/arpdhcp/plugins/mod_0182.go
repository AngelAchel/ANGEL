package arpdhcp

import (
    "time"
)

type arpdhcp0182 struct{}

func Newarpdhcp0182() *arpdhcp0182 {
    return &arpdhcp0182{}
}

func (e *arpdhcp0182) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0182) Name() string { return "arpdhcp0182" }
func (e *arpdhcp0182) Timestamp() time.Time { return time.Now() }
