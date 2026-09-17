package passwordreset

import (
    "time"
)

type passwordreset0130 struct{}

func Newpasswordreset0130() *passwordreset0130 {
    return &passwordreset0130{}
}

func (e *passwordreset0130) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0130) Name() string { return "passwordreset0130" }
func (e *passwordreset0130) Timestamp() time.Time { return time.Now() }
