package memcorrupt

import (
    "time"
)

type memcorrupt0083 struct{}

func Newmemcorrupt0083() *memcorrupt0083 {
    return &memcorrupt0083{}
}

func (e *memcorrupt0083) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0083) Name() string { return "memcorrupt0083" }
func (e *memcorrupt0083) Timestamp() time.Time { return time.Now() }
