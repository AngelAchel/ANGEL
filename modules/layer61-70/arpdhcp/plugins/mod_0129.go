package arpdhcp

import (
    "time"
)

type arpdhcp0129 struct{}

func Newarpdhcp0129() *arpdhcp0129 {
    return &arpdhcp0129{}
}

func (e *arpdhcp0129) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "arpdhcp:done")
    return results, nil
}

func (e *arpdhcp0129) Name() string { return "arpdhcp0129" }
func (e *arpdhcp0129) Timestamp() time.Time { return time.Now() }
