package passwordreset

import (
    "time"
)

type passwordreset0144 struct{}

func Newpasswordreset0144() *passwordreset0144 {
    return &passwordreset0144{}
}

func (e *passwordreset0144) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0144) Name() string { return "passwordreset0144" }
func (e *passwordreset0144) Timestamp() time.Time { return time.Now() }
