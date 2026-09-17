package passwordreset

import (
    "time"
)

type passwordreset0009 struct{}

func Newpasswordreset0009() *passwordreset0009 {
    return &passwordreset0009{}
}

func (e *passwordreset0009) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0009) Name() string { return "passwordreset0009" }
func (e *passwordreset0009) Timestamp() time.Time { return time.Now() }
