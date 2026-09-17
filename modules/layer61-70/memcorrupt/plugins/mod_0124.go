package memcorrupt

import (
    "time"
)

type memcorrupt0124 struct{}

func Newmemcorrupt0124() *memcorrupt0124 {
    return &memcorrupt0124{}
}

func (e *memcorrupt0124) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0124) Name() string { return "memcorrupt0124" }
func (e *memcorrupt0124) Timestamp() time.Time { return time.Now() }
