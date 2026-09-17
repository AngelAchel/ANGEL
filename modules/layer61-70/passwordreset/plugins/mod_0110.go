package passwordreset

import (
    "time"
)

type passwordreset0110 struct{}

func Newpasswordreset0110() *passwordreset0110 {
    return &passwordreset0110{}
}

func (e *passwordreset0110) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0110) Name() string { return "passwordreset0110" }
func (e *passwordreset0110) Timestamp() time.Time { return time.Now() }
