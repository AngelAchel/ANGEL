package memcorrupt

import (
    "time"
)

type memcorrupt0110 struct{}

func Newmemcorrupt0110() *memcorrupt0110 {
    return &memcorrupt0110{}
}

func (e *memcorrupt0110) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0110) Name() string { return "memcorrupt0110" }
func (e *memcorrupt0110) Timestamp() time.Time { return time.Now() }
