package passwordreset

import (
    "time"
)

type passwordreset0071 struct{}

func Newpasswordreset0071() *passwordreset0071 {
    return &passwordreset0071{}
}

func (e *passwordreset0071) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0071) Name() string { return "passwordreset0071" }
func (e *passwordreset0071) Timestamp() time.Time { return time.Now() }
