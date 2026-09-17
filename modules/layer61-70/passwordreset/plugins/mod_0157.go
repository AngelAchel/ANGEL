package passwordreset

import (
    "time"
)

type passwordreset0157 struct{}

func Newpasswordreset0157() *passwordreset0157 {
    return &passwordreset0157{}
}

func (e *passwordreset0157) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0157) Name() string { return "passwordreset0157" }
func (e *passwordreset0157) Timestamp() time.Time { return time.Now() }
