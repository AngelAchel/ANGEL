package memcorrupt

import (
    "time"
)

type memcorrupt0103 struct{}

func Newmemcorrupt0103() *memcorrupt0103 {
    return &memcorrupt0103{}
}

func (e *memcorrupt0103) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0103) Name() string { return "memcorrupt0103" }
func (e *memcorrupt0103) Timestamp() time.Time { return time.Now() }
