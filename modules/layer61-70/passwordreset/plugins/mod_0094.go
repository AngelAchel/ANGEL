package passwordreset

import (
    "time"
)

type passwordreset0094 struct{}

func Newpasswordreset0094() *passwordreset0094 {
    return &passwordreset0094{}
}

func (e *passwordreset0094) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0094) Name() string { return "passwordreset0094" }
func (e *passwordreset0094) Timestamp() time.Time { return time.Now() }
