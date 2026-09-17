package passwordreset

import (
    "time"
)

type passwordreset0093 struct{}

func Newpasswordreset0093() *passwordreset0093 {
    return &passwordreset0093{}
}

func (e *passwordreset0093) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0093) Name() string { return "passwordreset0093" }
func (e *passwordreset0093) Timestamp() time.Time { return time.Now() }
