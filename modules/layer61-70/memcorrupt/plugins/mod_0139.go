package memcorrupt

import (
    "time"
)

type memcorrupt0139 struct{}

func Newmemcorrupt0139() *memcorrupt0139 {
    return &memcorrupt0139{}
}

func (e *memcorrupt0139) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0139) Name() string { return "memcorrupt0139" }
func (e *memcorrupt0139) Timestamp() time.Time { return time.Now() }
