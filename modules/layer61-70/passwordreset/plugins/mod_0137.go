package passwordreset

import (
    "time"
)

type passwordreset0137 struct{}

func Newpasswordreset0137() *passwordreset0137 {
    return &passwordreset0137{}
}

func (e *passwordreset0137) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0137) Name() string { return "passwordreset0137" }
func (e *passwordreset0137) Timestamp() time.Time { return time.Now() }
