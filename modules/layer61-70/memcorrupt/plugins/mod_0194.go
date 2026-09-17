package memcorrupt

import (
    "time"
)

type memcorrupt0194 struct{}

func Newmemcorrupt0194() *memcorrupt0194 {
    return &memcorrupt0194{}
}

func (e *memcorrupt0194) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0194) Name() string { return "memcorrupt0194" }
func (e *memcorrupt0194) Timestamp() time.Time { return time.Now() }
