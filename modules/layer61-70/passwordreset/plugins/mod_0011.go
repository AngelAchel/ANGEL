package passwordreset

import (
    "time"
)

type passwordreset0011 struct{}

func Newpasswordreset0011() *passwordreset0011 {
    return &passwordreset0011{}
}

func (e *passwordreset0011) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0011) Name() string { return "passwordreset0011" }
func (e *passwordreset0011) Timestamp() time.Time { return time.Now() }
