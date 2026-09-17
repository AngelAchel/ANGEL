package deser

import (
    "time"
)

type deser0105 struct{}

func Newdeser0105() *deser0105 {
    return &deser0105{}
}

func (e *deser0105) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0105) Name() string { return "deser0105" }
func (e *deser0105) Timestamp() time.Time { return time.Now() }
