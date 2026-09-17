package passwordreset

import (
    "time"
)

type passwordreset0085 struct{}

func Newpasswordreset0085() *passwordreset0085 {
    return &passwordreset0085{}
}

func (e *passwordreset0085) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0085) Name() string { return "passwordreset0085" }
func (e *passwordreset0085) Timestamp() time.Time { return time.Now() }
