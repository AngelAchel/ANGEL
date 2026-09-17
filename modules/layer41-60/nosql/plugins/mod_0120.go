package nosql

import (
    "time"
)

type nosql0120 struct{}

func Newnosql0120() *nosql0120 {
    return &nosql0120{}
}

func (e *nosql0120) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "nosql:done")
    return results, nil
}

func (e *nosql0120) Name() string { return "nosql0120" }
func (e *nosql0120) Timestamp() time.Time { return time.Now() }
