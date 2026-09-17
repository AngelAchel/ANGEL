package arpdhcp

import (
    "time"
)

type arpdhcp0105 struct{}

func Newarpdhcp0105() *arpdhcp0105 {
    return &arpdhcp0105{}
}

func (e *arpdhcp0105) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0105) Name() string { return "arpdhcp0105" }
func (e *arpdhcp0105) Timestamp() time.Time { return time.Now() }
