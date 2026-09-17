package memcorrupt

import (
    "time"
)

type memcorrupt0190 struct{}

func Newmemcorrupt0190() *memcorrupt0190 {
    return &memcorrupt0190{}
}

func (e *memcorrupt0190) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0190) Name() string { return "memcorrupt0190" }
func (e *memcorrupt0190) Timestamp() time.Time { return time.Now() }
