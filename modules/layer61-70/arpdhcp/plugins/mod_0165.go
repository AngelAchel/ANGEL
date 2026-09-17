package arpdhcp

import (
    "time"
)

type arpdhcp0165 struct{}

func Newarpdhcp0165() *arpdhcp0165 {
    return &arpdhcp0165{}
}

func (e *arpdhcp0165) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0165) Name() string { return "arpdhcp0165" }
func (e *arpdhcp0165) Timestamp() time.Time { return time.Now() }
