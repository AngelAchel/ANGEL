package passwordreset

import (
    "time"
)

type passwordreset0012 struct{}

func Newpasswordreset0012() *passwordreset0012 {
    return &passwordreset0012{}
}

func (e *passwordreset0012) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0012) Name() string { return "passwordreset0012" }
func (e *passwordreset0012) Timestamp() time.Time { return time.Now() }
