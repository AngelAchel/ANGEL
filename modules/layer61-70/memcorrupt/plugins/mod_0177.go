package memcorrupt

import (
    "time"
)

type memcorrupt0177 struct{}

func Newmemcorrupt0177() *memcorrupt0177 {
    return &memcorrupt0177{}
}

func (e *memcorrupt0177) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0177) Name() string { return "memcorrupt0177" }
func (e *memcorrupt0177) Timestamp() time.Time { return time.Now() }
