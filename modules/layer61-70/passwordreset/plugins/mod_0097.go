package passwordreset

import (
    "time"
)

type passwordreset0097 struct{}

func Newpasswordreset0097() *passwordreset0097 {
    return &passwordreset0097{}
}

func (e *passwordreset0097) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0097) Name() string { return "passwordreset0097" }
func (e *passwordreset0097) Timestamp() time.Time { return time.Now() }
