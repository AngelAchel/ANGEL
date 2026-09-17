package passwordreset

import (
    "time"
)

type passwordreset0189 struct{}

func Newpasswordreset0189() *passwordreset0189 {
    return &passwordreset0189{}
}

func (e *passwordreset0189) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0189) Name() string { return "passwordreset0189" }
func (e *passwordreset0189) Timestamp() time.Time { return time.Now() }
