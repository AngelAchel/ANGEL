package passwordreset

import (
    "time"
)

type passwordreset0115 struct{}

func Newpasswordreset0115() *passwordreset0115 {
    return &passwordreset0115{}
}

func (e *passwordreset0115) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0115) Name() string { return "passwordreset0115" }
func (e *passwordreset0115) Timestamp() time.Time { return time.Now() }
