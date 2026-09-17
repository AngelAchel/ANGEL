package memcorrupt

import (
    "time"
)

type memcorrupt0112 struct{}

func Newmemcorrupt0112() *memcorrupt0112 {
    return &memcorrupt0112{}
}

func (e *memcorrupt0112) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0112) Name() string { return "memcorrupt0112" }
func (e *memcorrupt0112) Timestamp() time.Time { return time.Now() }
