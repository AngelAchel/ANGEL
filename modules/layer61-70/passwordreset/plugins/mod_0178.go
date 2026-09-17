package passwordreset

import (
    "time"
)

type passwordreset0178 struct{}

func Newpasswordreset0178() *passwordreset0178 {
    return &passwordreset0178{}
}

func (e *passwordreset0178) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0178) Name() string { return "passwordreset0178" }
func (e *passwordreset0178) Timestamp() time.Time { return time.Now() }
