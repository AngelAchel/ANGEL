package passwordreset

import (
    "time"
)

type passwordreset0139 struct{}

func Newpasswordreset0139() *passwordreset0139 {
    return &passwordreset0139{}
}

func (e *passwordreset0139) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0139) Name() string { return "passwordreset0139" }
func (e *passwordreset0139) Timestamp() time.Time { return time.Now() }
