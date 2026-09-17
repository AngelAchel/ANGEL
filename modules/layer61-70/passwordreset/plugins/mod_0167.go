package passwordreset

import (
    "time"
)

type passwordreset0167 struct{}

func Newpasswordreset0167() *passwordreset0167 {
    return &passwordreset0167{}
}

func (e *passwordreset0167) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0167) Name() string { return "passwordreset0167" }
func (e *passwordreset0167) Timestamp() time.Time { return time.Now() }
