package passwordreset

import (
    "time"
)

type passwordreset0111 struct{}

func Newpasswordreset0111() *passwordreset0111 {
    return &passwordreset0111{}
}

func (e *passwordreset0111) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0111) Name() string { return "passwordreset0111" }
func (e *passwordreset0111) Timestamp() time.Time { return time.Now() }
