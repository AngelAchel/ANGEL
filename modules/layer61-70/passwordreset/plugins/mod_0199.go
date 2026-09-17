package passwordreset

import (
    "time"
)

type passwordreset0199 struct{}

func Newpasswordreset0199() *passwordreset0199 {
    return &passwordreset0199{}
}

func (e *passwordreset0199) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0199) Name() string { return "passwordreset0199" }
func (e *passwordreset0199) Timestamp() time.Time { return time.Now() }
