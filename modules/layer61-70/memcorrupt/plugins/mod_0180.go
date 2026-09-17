package memcorrupt

import (
    "time"
)

type memcorrupt0180 struct{}

func Newmemcorrupt0180() *memcorrupt0180 {
    return &memcorrupt0180{}
}

func (e *memcorrupt0180) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0180) Name() string { return "memcorrupt0180" }
func (e *memcorrupt0180) Timestamp() time.Time { return time.Now() }
