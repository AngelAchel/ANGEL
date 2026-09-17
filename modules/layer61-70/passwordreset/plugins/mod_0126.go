package passwordreset

import (
    "time"
)

type passwordreset0126 struct{}

func Newpasswordreset0126() *passwordreset0126 {
    return &passwordreset0126{}
}

func (e *passwordreset0126) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0126) Name() string { return "passwordreset0126" }
func (e *passwordreset0126) Timestamp() time.Time { return time.Now() }
