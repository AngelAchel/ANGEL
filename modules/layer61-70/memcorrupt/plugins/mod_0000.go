package memcorrupt

import (
    "time"
)

type memcorrupt0000 struct{}

func Newmemcorrupt0000() *memcorrupt0000 {
    return &memcorrupt0000{}
}

func (e *memcorrupt0000) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0000) Name() string { return "memcorrupt0000" }
func (e *memcorrupt0000) Timestamp() time.Time { return time.Now() }
