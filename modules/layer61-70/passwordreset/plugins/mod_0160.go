package passwordreset

import (
    "time"
)

type passwordreset0160 struct{}

func Newpasswordreset0160() *passwordreset0160 {
    return &passwordreset0160{}
}

func (e *passwordreset0160) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0160) Name() string { return "passwordreset0160" }
func (e *passwordreset0160) Timestamp() time.Time { return time.Now() }
