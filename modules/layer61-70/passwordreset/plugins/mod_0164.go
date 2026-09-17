package passwordreset

import (
    "time"
)

type passwordreset0164 struct{}

func Newpasswordreset0164() *passwordreset0164 {
    return &passwordreset0164{}
}

func (e *passwordreset0164) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0164) Name() string { return "passwordreset0164" }
func (e *passwordreset0164) Timestamp() time.Time { return time.Now() }
