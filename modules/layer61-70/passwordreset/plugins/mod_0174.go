package passwordreset

import (
    "time"
)

type passwordreset0174 struct{}

func Newpasswordreset0174() *passwordreset0174 {
    return &passwordreset0174{}
}

func (e *passwordreset0174) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0174) Name() string { return "passwordreset0174" }
func (e *passwordreset0174) Timestamp() time.Time { return time.Now() }
