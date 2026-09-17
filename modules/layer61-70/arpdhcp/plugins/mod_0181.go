package arpdhcp

import (
    "time"
)

type arpdhcp0181 struct{}

func Newarpdhcp0181() *arpdhcp0181 {
    return &arpdhcp0181{}
}

func (e *arpdhcp0181) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0181) Name() string { return "arpdhcp0181" }
func (e *arpdhcp0181) Timestamp() time.Time { return time.Now() }
