package passwordreset

import (
    "time"
)

type passwordreset0136 struct{}

func Newpasswordreset0136() *passwordreset0136 {
    return &passwordreset0136{}
}

func (e *passwordreset0136) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0136) Name() string { return "passwordreset0136" }
func (e *passwordreset0136) Timestamp() time.Time { return time.Now() }
