package arpdhcp

import (
    "time"
)

type arpdhcp0053 struct{}

func Newarpdhcp0053() *arpdhcp0053 {
    return &arpdhcp0053{}
}

func (e *arpdhcp0053) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0053) Name() string { return "arpdhcp0053" }
func (e *arpdhcp0053) Timestamp() time.Time { return time.Now() }
