package memcorrupt

import (
    "time"
)

type memcorrupt0036 struct{}

func Newmemcorrupt0036() *memcorrupt0036 {
    return &memcorrupt0036{}
}

func (e *memcorrupt0036) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0036) Name() string { return "memcorrupt0036" }
func (e *memcorrupt0036) Timestamp() time.Time { return time.Now() }
