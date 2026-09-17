package memcorrupt

import (
    "time"
)

type memcorrupt0010 struct{}

func Newmemcorrupt0010() *memcorrupt0010 {
    return &memcorrupt0010{}
}

func (e *memcorrupt0010) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0010) Name() string { return "memcorrupt0010" }
func (e *memcorrupt0010) Timestamp() time.Time { return time.Now() }
