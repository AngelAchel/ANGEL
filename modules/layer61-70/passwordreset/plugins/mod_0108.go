package passwordreset

import (
    "time"
)

type passwordreset0108 struct{}

func Newpasswordreset0108() *passwordreset0108 {
    return &passwordreset0108{}
}

func (e *passwordreset0108) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0108) Name() string { return "passwordreset0108" }
func (e *passwordreset0108) Timestamp() time.Time { return time.Now() }
