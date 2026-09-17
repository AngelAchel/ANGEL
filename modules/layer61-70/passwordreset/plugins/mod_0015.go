package passwordreset

import (
    "time"
)

type passwordreset0015 struct{}

func Newpasswordreset0015() *passwordreset0015 {
    return &passwordreset0015{}
}

func (e *passwordreset0015) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0015) Name() string { return "passwordreset0015" }
func (e *passwordreset0015) Timestamp() time.Time { return time.Now() }
