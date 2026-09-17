package passwordreset

import (
    "time"
)

type passwordreset0107 struct{}

func Newpasswordreset0107() *passwordreset0107 {
    return &passwordreset0107{}
}

func (e *passwordreset0107) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0107) Name() string { return "passwordreset0107" }
func (e *passwordreset0107) Timestamp() time.Time { return time.Now() }
