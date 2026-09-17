package passwordreset

import (
    "time"
)

type passwordreset0121 struct{}

func Newpasswordreset0121() *passwordreset0121 {
    return &passwordreset0121{}
}

func (e *passwordreset0121) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0121) Name() string { return "passwordreset0121" }
func (e *passwordreset0121) Timestamp() time.Time { return time.Now() }
