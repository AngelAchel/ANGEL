package memcorrupt

import (
    "time"
)

type memcorrupt0117 struct{}

func Newmemcorrupt0117() *memcorrupt0117 {
    return &memcorrupt0117{}
}

func (e *memcorrupt0117) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0117) Name() string { return "memcorrupt0117" }
func (e *memcorrupt0117) Timestamp() time.Time { return time.Now() }
