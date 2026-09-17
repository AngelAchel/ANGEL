package passwordreset

import (
    "time"
)

type passwordreset0104 struct{}

func Newpasswordreset0104() *passwordreset0104 {
    return &passwordreset0104{}
}

func (e *passwordreset0104) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0104) Name() string { return "passwordreset0104" }
func (e *passwordreset0104) Timestamp() time.Time { return time.Now() }
