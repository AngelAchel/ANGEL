package memcorrupt

import (
    "time"
)

type memcorrupt0120 struct{}

func Newmemcorrupt0120() *memcorrupt0120 {
    return &memcorrupt0120{}
}

func (e *memcorrupt0120) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0120) Name() string { return "memcorrupt0120" }
func (e *memcorrupt0120) Timestamp() time.Time { return time.Now() }
