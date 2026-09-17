package passwordreset

import (
    "time"
)

type passwordreset0153 struct{}

func Newpasswordreset0153() *passwordreset0153 {
    return &passwordreset0153{}
}

func (e *passwordreset0153) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0153) Name() string { return "passwordreset0153" }
func (e *passwordreset0153) Timestamp() time.Time { return time.Now() }
