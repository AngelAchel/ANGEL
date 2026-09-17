package memcorrupt

import (
    "time"
)

type memcorrupt0105 struct{}

func Newmemcorrupt0105() *memcorrupt0105 {
    return &memcorrupt0105{}
}

func (e *memcorrupt0105) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0105) Name() string { return "memcorrupt0105" }
func (e *memcorrupt0105) Timestamp() time.Time { return time.Now() }
