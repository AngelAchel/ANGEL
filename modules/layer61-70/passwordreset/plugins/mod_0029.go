package passwordreset

import (
    "time"
)

type passwordreset0029 struct{}

func Newpasswordreset0029() *passwordreset0029 {
    return &passwordreset0029{}
}

func (e *passwordreset0029) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0029) Name() string { return "passwordreset0029" }
func (e *passwordreset0029) Timestamp() time.Time { return time.Now() }
