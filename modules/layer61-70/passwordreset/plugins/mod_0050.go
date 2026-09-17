package passwordreset

import (
    "time"
)

type passwordreset0050 struct{}

func Newpasswordreset0050() *passwordreset0050 {
    return &passwordreset0050{}
}

func (e *passwordreset0050) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0050) Name() string { return "passwordreset0050" }
func (e *passwordreset0050) Timestamp() time.Time { return time.Now() }
