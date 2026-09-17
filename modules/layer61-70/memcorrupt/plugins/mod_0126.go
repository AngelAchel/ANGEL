package memcorrupt

import (
    "time"
)

type memcorrupt0126 struct{}

func Newmemcorrupt0126() *memcorrupt0126 {
    return &memcorrupt0126{}
}

func (e *memcorrupt0126) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0126) Name() string { return "memcorrupt0126" }
func (e *memcorrupt0126) Timestamp() time.Time { return time.Now() }
