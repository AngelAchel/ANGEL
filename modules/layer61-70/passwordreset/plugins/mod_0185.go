package passwordreset

import (
    "time"
)

type passwordreset0185 struct{}

func Newpasswordreset0185() *passwordreset0185 {
    return &passwordreset0185{}
}

func (e *passwordreset0185) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0185) Name() string { return "passwordreset0185" }
func (e *passwordreset0185) Timestamp() time.Time { return time.Now() }
