package memcorrupt

import (
    "time"
)

type memcorrupt0197 struct{}

func Newmemcorrupt0197() *memcorrupt0197 {
    return &memcorrupt0197{}
}

func (e *memcorrupt0197) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0197) Name() string { return "memcorrupt0197" }
func (e *memcorrupt0197) Timestamp() time.Time { return time.Now() }
