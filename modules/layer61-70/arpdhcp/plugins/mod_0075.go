package arpdhcp

import (
    "time"
)

type arpdhcp0075 struct{}

func Newarpdhcp0075() *arpdhcp0075 {
    return &arpdhcp0075{}
}

func (e *arpdhcp0075) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0075) Name() string { return "arpdhcp0075" }
func (e *arpdhcp0075) Timestamp() time.Time { return time.Now() }
