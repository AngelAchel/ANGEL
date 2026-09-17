package memcorrupt

import (
    "time"
)

type memcorrupt0101 struct{}

func Newmemcorrupt0101() *memcorrupt0101 {
    return &memcorrupt0101{}
}

func (e *memcorrupt0101) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0101) Name() string { return "memcorrupt0101" }
func (e *memcorrupt0101) Timestamp() time.Time { return time.Now() }
