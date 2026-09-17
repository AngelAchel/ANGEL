package passwordreset

import (
    "time"
)

type passwordreset0016 struct{}

func Newpasswordreset0016() *passwordreset0016 {
    return &passwordreset0016{}
}

func (e *passwordreset0016) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0016) Name() string { return "passwordreset0016" }
func (e *passwordreset0016) Timestamp() time.Time { return time.Now() }
