package passwordreset

import (
    "time"
)

type passwordreset0018 struct{}

func Newpasswordreset0018() *passwordreset0018 {
    return &passwordreset0018{}
}

func (e *passwordreset0018) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0018) Name() string { return "passwordreset0018" }
func (e *passwordreset0018) Timestamp() time.Time { return time.Now() }
