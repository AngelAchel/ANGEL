package passwordreset

import (
    "time"
)

type passwordreset0147 struct{}

func Newpasswordreset0147() *passwordreset0147 {
    return &passwordreset0147{}
}

func (e *passwordreset0147) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0147) Name() string { return "passwordreset0147" }
func (e *passwordreset0147) Timestamp() time.Time { return time.Now() }
