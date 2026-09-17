package passwordreset

import (
    "time"
)

type passwordreset0101 struct{}

func Newpasswordreset0101() *passwordreset0101 {
    return &passwordreset0101{}
}

func (e *passwordreset0101) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0101) Name() string { return "passwordreset0101" }
func (e *passwordreset0101) Timestamp() time.Time { return time.Now() }
