package memcorrupt

import (
    "time"
)

type memcorrupt0121 struct{}

func Newmemcorrupt0121() *memcorrupt0121 {
    return &memcorrupt0121{}
}

func (e *memcorrupt0121) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0121) Name() string { return "memcorrupt0121" }
func (e *memcorrupt0121) Timestamp() time.Time { return time.Now() }
