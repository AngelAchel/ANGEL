package passwordreset

import (
    "time"
)

type passwordreset0186 struct{}

func Newpasswordreset0186() *passwordreset0186 {
    return &passwordreset0186{}
}

func (e *passwordreset0186) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0186) Name() string { return "passwordreset0186" }
func (e *passwordreset0186) Timestamp() time.Time { return time.Now() }
