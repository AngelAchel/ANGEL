package passwordreset

import (
    "time"
)

type passwordreset0089 struct{}

func Newpasswordreset0089() *passwordreset0089 {
    return &passwordreset0089{}
}

func (e *passwordreset0089) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0089) Name() string { return "passwordreset0089" }
func (e *passwordreset0089) Timestamp() time.Time { return time.Now() }
