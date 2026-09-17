package passwordreset

import (
    "time"
)

type passwordreset0120 struct{}

func Newpasswordreset0120() *passwordreset0120 {
    return &passwordreset0120{}
}

func (e *passwordreset0120) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0120) Name() string { return "passwordreset0120" }
func (e *passwordreset0120) Timestamp() time.Time { return time.Now() }
