package nosql

import (
    "time"
)

type nosql0101 struct{}

func Newnosql0101() *nosql0101 {
    return &nosql0101{}
}

func (e *nosql0101) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "nosql:done")
    return results, nil
}

func (e *nosql0101) Name() string { return "nosql0101" }
func (e *nosql0101) Timestamp() time.Time { return time.Now() }
