package passwordreset

import (
    "time"
)

type passwordreset0010 struct{}

func Newpasswordreset0010() *passwordreset0010 {
    return &passwordreset0010{}
}

func (e *passwordreset0010) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0010) Name() string { return "passwordreset0010" }
func (e *passwordreset0010) Timestamp() time.Time { return time.Now() }
