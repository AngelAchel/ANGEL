package passwordreset

import (
    "time"
)

type passwordreset0041 struct{}

func Newpasswordreset0041() *passwordreset0041 {
    return &passwordreset0041{}
}

func (e *passwordreset0041) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0041) Name() string { return "passwordreset0041" }
func (e *passwordreset0041) Timestamp() time.Time { return time.Now() }
