package passwordreset

import (
    "time"
)

type passwordreset0175 struct{}

func Newpasswordreset0175() *passwordreset0175 {
    return &passwordreset0175{}
}

func (e *passwordreset0175) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0175) Name() string { return "passwordreset0175" }
func (e *passwordreset0175) Timestamp() time.Time { return time.Now() }
