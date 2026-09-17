package memcorrupt

import (
    "time"
)

type memcorrupt0123 struct{}

func Newmemcorrupt0123() *memcorrupt0123 {
    return &memcorrupt0123{}
}

func (e *memcorrupt0123) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0123) Name() string { return "memcorrupt0123" }
func (e *memcorrupt0123) Timestamp() time.Time { return time.Now() }
