package passwordreset

import (
    "time"
)

type passwordreset0030 struct{}

func Newpasswordreset0030() *passwordreset0030 {
    return &passwordreset0030{}
}

func (e *passwordreset0030) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0030) Name() string { return "passwordreset0030" }
func (e *passwordreset0030) Timestamp() time.Time { return time.Now() }
