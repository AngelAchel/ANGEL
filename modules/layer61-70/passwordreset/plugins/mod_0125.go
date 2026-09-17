package passwordreset

import (
    "time"
)

type passwordreset0125 struct{}

func Newpasswordreset0125() *passwordreset0125 {
    return &passwordreset0125{}
}

func (e *passwordreset0125) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0125) Name() string { return "passwordreset0125" }
func (e *passwordreset0125) Timestamp() time.Time { return time.Now() }
