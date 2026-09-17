package memcorrupt

import (
    "time"
)

type memcorrupt0173 struct{}

func Newmemcorrupt0173() *memcorrupt0173 {
    return &memcorrupt0173{}
}

func (e *memcorrupt0173) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0173) Name() string { return "memcorrupt0173" }
func (e *memcorrupt0173) Timestamp() time.Time { return time.Now() }
