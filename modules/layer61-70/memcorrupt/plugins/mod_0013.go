package memcorrupt

import (
    "time"
)

type memcorrupt0013 struct{}

func Newmemcorrupt0013() *memcorrupt0013 {
    return &memcorrupt0013{}
}

func (e *memcorrupt0013) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0013) Name() string { return "memcorrupt0013" }
func (e *memcorrupt0013) Timestamp() time.Time { return time.Now() }
