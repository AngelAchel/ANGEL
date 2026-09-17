package deser

import (
    "time"
)

type deser0038 struct{}

func Newdeser0038() *deser0038 {
    return &deser0038{}
}

func (e *deser0038) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0038) Name() string { return "deser0038" }
func (e *deser0038) Timestamp() time.Time { return time.Now() }
