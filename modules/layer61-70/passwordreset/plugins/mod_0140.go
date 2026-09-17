package passwordreset

import (
    "time"
)

type passwordreset0140 struct{}

func Newpasswordreset0140() *passwordreset0140 {
    return &passwordreset0140{}
}

func (e *passwordreset0140) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0140) Name() string { return "passwordreset0140" }
func (e *passwordreset0140) Timestamp() time.Time { return time.Now() }
