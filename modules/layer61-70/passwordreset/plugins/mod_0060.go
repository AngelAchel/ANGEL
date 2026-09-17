package passwordreset

import (
    "time"
)

type passwordreset0060 struct{}

func Newpasswordreset0060() *passwordreset0060 {
    return &passwordreset0060{}
}

func (e *passwordreset0060) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0060) Name() string { return "passwordreset0060" }
func (e *passwordreset0060) Timestamp() time.Time { return time.Now() }
