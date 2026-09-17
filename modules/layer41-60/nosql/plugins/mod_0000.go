package nosql

import (
    "time"
)

type nosql0000 struct{}

func Newnosql0000() *nosql0000 {
    return &nosql0000{}
}

func (e *nosql0000) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "nosql:done")
    return results, nil
}

func (e *nosql0000) Name() string { return "nosql0000" }
func (e *nosql0000) Timestamp() time.Time { return time.Now() }
