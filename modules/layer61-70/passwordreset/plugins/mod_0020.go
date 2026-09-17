package passwordreset

import (
    "time"
)

type passwordreset0020 struct{}

func Newpasswordreset0020() *passwordreset0020 {
    return &passwordreset0020{}
}

func (e *passwordreset0020) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0020) Name() string { return "passwordreset0020" }
func (e *passwordreset0020) Timestamp() time.Time { return time.Now() }
