package passwordreset

import (
    "time"
)

type passwordreset0013 struct{}

func Newpasswordreset0013() *passwordreset0013 {
    return &passwordreset0013{}
}

func (e *passwordreset0013) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0013) Name() string { return "passwordreset0013" }
func (e *passwordreset0013) Timestamp() time.Time { return time.Now() }
