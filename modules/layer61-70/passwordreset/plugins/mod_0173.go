package passwordreset

import (
    "time"
)

type passwordreset0173 struct{}

func Newpasswordreset0173() *passwordreset0173 {
    return &passwordreset0173{}
}

func (e *passwordreset0173) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0173) Name() string { return "passwordreset0173" }
func (e *passwordreset0173) Timestamp() time.Time { return time.Now() }
