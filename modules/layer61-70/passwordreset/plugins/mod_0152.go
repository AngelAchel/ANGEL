package passwordreset

import (
    "time"
)

type passwordreset0152 struct{}

func Newpasswordreset0152() *passwordreset0152 {
    return &passwordreset0152{}
}

func (e *passwordreset0152) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0152) Name() string { return "passwordreset0152" }
func (e *passwordreset0152) Timestamp() time.Time { return time.Now() }
