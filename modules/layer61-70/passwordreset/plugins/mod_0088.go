package passwordreset

import (
    "time"
)

type passwordreset0088 struct{}

func Newpasswordreset0088() *passwordreset0088 {
    return &passwordreset0088{}
}

func (e *passwordreset0088) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0088) Name() string { return "passwordreset0088" }
func (e *passwordreset0088) Timestamp() time.Time { return time.Now() }
