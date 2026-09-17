package passwordreset

import (
    "time"
)

type passwordreset0188 struct{}

func Newpasswordreset0188() *passwordreset0188 {
    return &passwordreset0188{}
}

func (e *passwordreset0188) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0188) Name() string { return "passwordreset0188" }
func (e *passwordreset0188) Timestamp() time.Time { return time.Now() }
