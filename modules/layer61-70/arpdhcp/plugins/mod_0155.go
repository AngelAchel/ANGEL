package arpdhcp

import (
    "time"
)

type arpdhcp0155 struct{}

func Newarpdhcp0155() *arpdhcp0155 {
    return &arpdhcp0155{}
}

func (e *arpdhcp0155) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0155) Name() string { return "arpdhcp0155" }
func (e *arpdhcp0155) Timestamp() time.Time { return time.Now() }
