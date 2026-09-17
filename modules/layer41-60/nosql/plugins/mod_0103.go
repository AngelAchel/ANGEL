package nosql

import (
    "time"
)

type nosql0103 struct{}

func Newnosql0103() *nosql0103 {
    return &nosql0103{}
}

func (e *nosql0103) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "nosql:done")
    return results, nil
}

func (e *nosql0103) Name() string { return "nosql0103" }
func (e *nosql0103) Timestamp() time.Time { return time.Now() }
