package memcorrupt

import (
    "time"
)

type memcorrupt0160 struct{}

func Newmemcorrupt0160() *memcorrupt0160 {
    return &memcorrupt0160{}
}

func (e *memcorrupt0160) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0160) Name() string { return "memcorrupt0160" }
func (e *memcorrupt0160) Timestamp() time.Time { return time.Now() }
