package arpdhcp

import (
    "time"
)

type arpdhcp0115 struct{}

func Newarpdhcp0115() *arpdhcp0115 {
    return &arpdhcp0115{}
}

func (e *arpdhcp0115) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0115) Name() string { return "arpdhcp0115" }
func (e *arpdhcp0115) Timestamp() time.Time { return time.Now() }
