package passwordreset

import (
    "time"
)

type passwordreset0025 struct{}

func Newpasswordreset0025() *passwordreset0025 {
    return &passwordreset0025{}
}

func (e *passwordreset0025) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0025) Name() string { return "passwordreset0025" }
func (e *passwordreset0025) Timestamp() time.Time { return time.Now() }
