package arpdhcp

import (
    "time"
)

type arpdhcp0184 struct{}

func Newarpdhcp0184() *arpdhcp0184 {
    return &arpdhcp0184{}
}

func (e *arpdhcp0184) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0184) Name() string { return "arpdhcp0184" }
func (e *arpdhcp0184) Timestamp() time.Time { return time.Now() }
