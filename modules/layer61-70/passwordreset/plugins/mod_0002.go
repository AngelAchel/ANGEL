package passwordreset

import (
    "time"
)

type passwordreset0002 struct{}

func Newpasswordreset0002() *passwordreset0002 {
    return &passwordreset0002{}
}

func (e *passwordreset0002) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0002) Name() string { return "passwordreset0002" }
func (e *passwordreset0002) Timestamp() time.Time { return time.Now() }
