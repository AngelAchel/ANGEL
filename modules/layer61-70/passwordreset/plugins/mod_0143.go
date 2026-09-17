package passwordreset

import (
    "time"
)

type passwordreset0143 struct{}

func Newpasswordreset0143() *passwordreset0143 {
    return &passwordreset0143{}
}

func (e *passwordreset0143) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0143) Name() string { return "passwordreset0143" }
func (e *passwordreset0143) Timestamp() time.Time { return time.Now() }
