package passwordreset

import (
    "time"
)

type passwordreset0006 struct{}

func Newpasswordreset0006() *passwordreset0006 {
    return &passwordreset0006{}
}

func (e *passwordreset0006) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0006) Name() string { return "passwordreset0006" }
func (e *passwordreset0006) Timestamp() time.Time { return time.Now() }
