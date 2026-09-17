package memcorrupt

import (
    "time"
)

type memcorrupt0172 struct{}

func Newmemcorrupt0172() *memcorrupt0172 {
    return &memcorrupt0172{}
}

func (e *memcorrupt0172) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0172) Name() string { return "memcorrupt0172" }
func (e *memcorrupt0172) Timestamp() time.Time { return time.Now() }
