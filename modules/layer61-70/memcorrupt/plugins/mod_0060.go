package memcorrupt

import (
    "time"
)

type memcorrupt0060 struct{}

func Newmemcorrupt0060() *memcorrupt0060 {
    return &memcorrupt0060{}
}

func (e *memcorrupt0060) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0060) Name() string { return "memcorrupt0060" }
func (e *memcorrupt0060) Timestamp() time.Time { return time.Now() }
