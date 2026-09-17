package passwordreset

import (
    "time"
)

type passwordreset0037 struct{}

func Newpasswordreset0037() *passwordreset0037 {
    return &passwordreset0037{}
}

func (e *passwordreset0037) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0037) Name() string { return "passwordreset0037" }
func (e *passwordreset0037) Timestamp() time.Time { return time.Now() }
