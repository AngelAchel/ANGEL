package deser

import (
    "time"
)

type deser0171 struct{}

func Newdeser0171() *deser0171 {
    return &deser0171{}
}

func (e *deser0171) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0171) Name() string { return "deser0171" }
func (e *deser0171) Timestamp() time.Time { return time.Now() }
