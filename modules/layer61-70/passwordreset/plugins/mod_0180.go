package passwordreset

import (
    "time"
)

type passwordreset0180 struct{}

func Newpasswordreset0180() *passwordreset0180 {
    return &passwordreset0180{}
}

func (e *passwordreset0180) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0180) Name() string { return "passwordreset0180" }
func (e *passwordreset0180) Timestamp() time.Time { return time.Now() }
