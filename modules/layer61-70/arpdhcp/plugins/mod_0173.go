package arpdhcp

import (
    "time"
)

type arpdhcp0173 struct{}

func Newarpdhcp0173() *arpdhcp0173 {
    return &arpdhcp0173{}
}

func (e *arpdhcp0173) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0173) Name() string { return "arpdhcp0173" }
func (e *arpdhcp0173) Timestamp() time.Time { return time.Now() }
