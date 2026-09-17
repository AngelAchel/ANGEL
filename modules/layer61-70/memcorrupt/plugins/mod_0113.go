package memcorrupt

import (
    "time"
)

type memcorrupt0113 struct{}

func Newmemcorrupt0113() *memcorrupt0113 {
    return &memcorrupt0113{}
}

func (e *memcorrupt0113) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0113) Name() string { return "memcorrupt0113" }
func (e *memcorrupt0113) Timestamp() time.Time { return time.Now() }
