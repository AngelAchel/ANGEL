package passwordreset

import (
    "time"
)

type passwordreset0086 struct{}

func Newpasswordreset0086() *passwordreset0086 {
    return &passwordreset0086{}
}

func (e *passwordreset0086) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0086) Name() string { return "passwordreset0086" }
func (e *passwordreset0086) Timestamp() time.Time { return time.Now() }
