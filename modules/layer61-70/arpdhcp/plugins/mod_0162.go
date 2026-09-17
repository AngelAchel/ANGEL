package arpdhcp

import (
    "time"
)

type arpdhcp0162 struct{}

func Newarpdhcp0162() *arpdhcp0162 {
    return &arpdhcp0162{}
}

func (e *arpdhcp0162) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0162) Name() string { return "arpdhcp0162" }
func (e *arpdhcp0162) Timestamp() time.Time { return time.Now() }
