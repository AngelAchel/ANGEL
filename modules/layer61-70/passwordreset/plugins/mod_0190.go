package passwordreset

import (
    "time"
)

type passwordreset0190 struct{}

func Newpasswordreset0190() *passwordreset0190 {
    return &passwordreset0190{}
}

func (e *passwordreset0190) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0190) Name() string { return "passwordreset0190" }
func (e *passwordreset0190) Timestamp() time.Time { return time.Now() }
