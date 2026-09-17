package passwordreset

import (
    "time"
)

type passwordreset0072 struct{}

func Newpasswordreset0072() *passwordreset0072 {
    return &passwordreset0072{}
}

func (e *passwordreset0072) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0072) Name() string { return "passwordreset0072" }
func (e *passwordreset0072) Timestamp() time.Time { return time.Now() }
