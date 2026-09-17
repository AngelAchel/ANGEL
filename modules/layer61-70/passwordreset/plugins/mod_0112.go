package passwordreset

import (
    "time"
)

type passwordreset0112 struct{}

func Newpasswordreset0112() *passwordreset0112 {
    return &passwordreset0112{}
}

func (e *passwordreset0112) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0112) Name() string { return "passwordreset0112" }
func (e *passwordreset0112) Timestamp() time.Time { return time.Now() }
