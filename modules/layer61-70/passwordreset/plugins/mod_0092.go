package passwordreset

import (
    "time"
)

type passwordreset0092 struct{}

func Newpasswordreset0092() *passwordreset0092 {
    return &passwordreset0092{}
}

func (e *passwordreset0092) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0092) Name() string { return "passwordreset0092" }
func (e *passwordreset0092) Timestamp() time.Time { return time.Now() }
