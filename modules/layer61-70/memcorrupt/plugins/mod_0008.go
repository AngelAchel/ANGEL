package memcorrupt

import (
    "time"
)

type memcorrupt0008 struct{}

func Newmemcorrupt0008() *memcorrupt0008 {
    return &memcorrupt0008{}
}

func (e *memcorrupt0008) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0008) Name() string { return "memcorrupt0008" }
func (e *memcorrupt0008) Timestamp() time.Time { return time.Now() }
