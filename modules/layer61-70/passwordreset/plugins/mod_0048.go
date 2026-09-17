package passwordreset

import (
    "time"
)

type passwordreset0048 struct{}

func Newpasswordreset0048() *passwordreset0048 {
    return &passwordreset0048{}
}

func (e *passwordreset0048) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0048) Name() string { return "passwordreset0048" }
func (e *passwordreset0048) Timestamp() time.Time { return time.Now() }
