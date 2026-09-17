package memcorrupt

import (
    "time"
)

type memcorrupt0100 struct{}

func Newmemcorrupt0100() *memcorrupt0100 {
    return &memcorrupt0100{}
}

func (e *memcorrupt0100) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0100) Name() string { return "memcorrupt0100" }
func (e *memcorrupt0100) Timestamp() time.Time { return time.Now() }
