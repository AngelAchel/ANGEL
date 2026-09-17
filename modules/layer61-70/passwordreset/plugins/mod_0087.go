package passwordreset

import (
    "time"
)

type passwordreset0087 struct{}

func Newpasswordreset0087() *passwordreset0087 {
    return &passwordreset0087{}
}

func (e *passwordreset0087) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0087) Name() string { return "passwordreset0087" }
func (e *passwordreset0087) Timestamp() time.Time { return time.Now() }
