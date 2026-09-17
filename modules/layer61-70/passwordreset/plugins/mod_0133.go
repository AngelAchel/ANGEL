package passwordreset

import (
    "time"
)

type passwordreset0133 struct{}

func Newpasswordreset0133() *passwordreset0133 {
    return &passwordreset0133{}
}

func (e *passwordreset0133) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0133) Name() string { return "passwordreset0133" }
func (e *passwordreset0133) Timestamp() time.Time { return time.Now() }
