package passwordreset

import (
    "time"
)

type passwordreset0123 struct{}

func Newpasswordreset0123() *passwordreset0123 {
    return &passwordreset0123{}
}

func (e *passwordreset0123) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0123) Name() string { return "passwordreset0123" }
func (e *passwordreset0123) Timestamp() time.Time { return time.Now() }
