package passwordreset

import (
    "time"
)

type passwordreset0105 struct{}

func Newpasswordreset0105() *passwordreset0105 {
    return &passwordreset0105{}
}

func (e *passwordreset0105) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0105) Name() string { return "passwordreset0105" }
func (e *passwordreset0105) Timestamp() time.Time { return time.Now() }
