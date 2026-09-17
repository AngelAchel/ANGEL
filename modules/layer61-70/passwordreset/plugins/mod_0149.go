package passwordreset

import (
    "time"
)

type passwordreset0149 struct{}

func Newpasswordreset0149() *passwordreset0149 {
    return &passwordreset0149{}
}

func (e *passwordreset0149) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0149) Name() string { return "passwordreset0149" }
func (e *passwordreset0149) Timestamp() time.Time { return time.Now() }
