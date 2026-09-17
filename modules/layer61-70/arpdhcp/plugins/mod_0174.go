package arpdhcp

import (
    "time"
)

type arpdhcp0174 struct{}

func Newarpdhcp0174() *arpdhcp0174 {
    return &arpdhcp0174{}
}

func (e *arpdhcp0174) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0174) Name() string { return "arpdhcp0174" }
func (e *arpdhcp0174) Timestamp() time.Time { return time.Now() }
