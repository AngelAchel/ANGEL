package memcorrupt

import (
    "time"
)

type memcorrupt0082 struct{}

func Newmemcorrupt0082() *memcorrupt0082 {
    return &memcorrupt0082{}
}

func (e *memcorrupt0082) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0082) Name() string { return "memcorrupt0082" }
func (e *memcorrupt0082) Timestamp() time.Time { return time.Now() }
