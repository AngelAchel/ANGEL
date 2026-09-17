package arpdhcp

import (
    "time"
)

type arpdhcp0057 struct{}

func Newarpdhcp0057() *arpdhcp0057 {
    return &arpdhcp0057{}
}

func (e *arpdhcp0057) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0057) Name() string { return "arpdhcp0057" }
func (e *arpdhcp0057) Timestamp() time.Time { return time.Now() }
