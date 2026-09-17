package passwordreset

import (
    "time"
)

type passwordreset0172 struct{}

func Newpasswordreset0172() *passwordreset0172 {
    return &passwordreset0172{}
}

func (e *passwordreset0172) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0172) Name() string { return "passwordreset0172" }
func (e *passwordreset0172) Timestamp() time.Time { return time.Now() }
