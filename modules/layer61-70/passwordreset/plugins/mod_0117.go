package passwordreset

import (
    "time"
)

type passwordreset0117 struct{}

func Newpasswordreset0117() *passwordreset0117 {
    return &passwordreset0117{}
}

func (e *passwordreset0117) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0117) Name() string { return "passwordreset0117" }
func (e *passwordreset0117) Timestamp() time.Time { return time.Now() }
