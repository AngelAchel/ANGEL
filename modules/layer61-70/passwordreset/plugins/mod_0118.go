package passwordreset

import (
    "time"
)

type passwordreset0118 struct{}

func Newpasswordreset0118() *passwordreset0118 {
    return &passwordreset0118{}
}

func (e *passwordreset0118) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0118) Name() string { return "passwordreset0118" }
func (e *passwordreset0118) Timestamp() time.Time { return time.Now() }
