package passwordreset

import (
    "time"
)

type passwordreset0142 struct{}

func Newpasswordreset0142() *passwordreset0142 {
    return &passwordreset0142{}
}

func (e *passwordreset0142) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0142) Name() string { return "passwordreset0142" }
func (e *passwordreset0142) Timestamp() time.Time { return time.Now() }
