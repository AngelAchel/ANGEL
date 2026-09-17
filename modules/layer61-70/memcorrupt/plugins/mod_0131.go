package memcorrupt

import (
    "time"
)

type memcorrupt0131 struct{}

func Newmemcorrupt0131() *memcorrupt0131 {
    return &memcorrupt0131{}
}

func (e *memcorrupt0131) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "memcorrupt:done")
    return results, nil
}

func (e *memcorrupt0131) Name() string { return "memcorrupt0131" }
func (e *memcorrupt0131) Timestamp() time.Time { return time.Now() }
