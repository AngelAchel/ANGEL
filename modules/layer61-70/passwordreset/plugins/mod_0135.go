package passwordreset

import (
    "time"
)

type passwordreset0135 struct{}

func Newpasswordreset0135() *passwordreset0135 {
    return &passwordreset0135{}
}

func (e *passwordreset0135) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0135) Name() string { return "passwordreset0135" }
func (e *passwordreset0135) Timestamp() time.Time { return time.Now() }
