package passwordreset

import (
    "time"
)

type passwordreset0003 struct{}

func Newpasswordreset0003() *passwordreset0003 {
    return &passwordreset0003{}
}

func (e *passwordreset0003) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0003) Name() string { return "passwordreset0003" }
func (e *passwordreset0003) Timestamp() time.Time { return time.Now() }
