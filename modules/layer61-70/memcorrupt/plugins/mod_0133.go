package memcorrupt

import (
    "time"
)

type memcorrupt0133 struct{}

func Newmemcorrupt0133() *memcorrupt0133 {
    return &memcorrupt0133{}
}

func (e *memcorrupt0133) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0133) Name() string { return "memcorrupt0133" }
func (e *memcorrupt0133) Timestamp() time.Time { return time.Now() }
