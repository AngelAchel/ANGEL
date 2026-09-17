package memcorrupt

import (
    "time"
)

type memcorrupt0018 struct{}

func Newmemcorrupt0018() *memcorrupt0018 {
    return &memcorrupt0018{}
}

func (e *memcorrupt0018) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0018) Name() string { return "memcorrupt0018" }
func (e *memcorrupt0018) Timestamp() time.Time { return time.Now() }
