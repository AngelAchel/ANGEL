package deser

import (
    "time"
)

type deser0173 struct{}

func Newdeser0173() *deser0173 {
    return &deser0173{}
}

func (e *deser0173) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "deser:done")
    return results, nil
}

func (e *deser0173) Name() string { return "deser0173" }
func (e *deser0173) Timestamp() time.Time { return time.Now() }
