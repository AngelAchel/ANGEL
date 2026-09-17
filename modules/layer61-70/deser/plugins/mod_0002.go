package deser

import (
    "time"
)

type deser0002 struct{}

func Newdeser0002() *deser0002 {
    return &deser0002{}
}

func (e *deser0002) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0002) Name() string { return "deser0002" }
func (e *deser0002) Timestamp() time.Time { return time.Now() }
