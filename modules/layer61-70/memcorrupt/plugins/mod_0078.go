package memcorrupt

import (
    "time"
)

type memcorrupt0078 struct{}

func Newmemcorrupt0078() *memcorrupt0078 {
    return &memcorrupt0078{}
}

func (e *memcorrupt0078) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0078) Name() string { return "memcorrupt0078" }
func (e *memcorrupt0078) Timestamp() time.Time { return time.Now() }
