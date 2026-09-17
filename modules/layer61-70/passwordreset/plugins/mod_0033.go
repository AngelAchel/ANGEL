package passwordreset

import (
    "time"
)

type passwordreset0033 struct{}

func Newpasswordreset0033() *passwordreset0033 {
    return &passwordreset0033{}
}

func (e *passwordreset0033) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0033) Name() string { return "passwordreset0033" }
func (e *passwordreset0033) Timestamp() time.Time { return time.Now() }
