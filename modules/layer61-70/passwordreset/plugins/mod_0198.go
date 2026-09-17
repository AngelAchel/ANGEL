package passwordreset

import (
    "time"
)

type passwordreset0198 struct{}

func Newpasswordreset0198() *passwordreset0198 {
    return &passwordreset0198{}
}

func (e *passwordreset0198) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0198) Name() string { return "passwordreset0198" }
func (e *passwordreset0198) Timestamp() time.Time { return time.Now() }
