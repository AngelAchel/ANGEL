package memcorrupt

import (
    "time"
)

type memcorrupt0037 struct{}

func Newmemcorrupt0037() *memcorrupt0037 {
    return &memcorrupt0037{}
}

func (e *memcorrupt0037) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0037) Name() string { return "memcorrupt0037" }
func (e *memcorrupt0037) Timestamp() time.Time { return time.Now() }
