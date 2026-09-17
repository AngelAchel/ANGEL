package memcorrupt

import (
    "time"
)

type memcorrupt0118 struct{}

func Newmemcorrupt0118() *memcorrupt0118 {
    return &memcorrupt0118{}
}

func (e *memcorrupt0118) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0118) Name() string { return "memcorrupt0118" }
func (e *memcorrupt0118) Timestamp() time.Time { return time.Now() }
