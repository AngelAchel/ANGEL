package passwordreset

import (
    "time"
)

type passwordreset0040 struct{}

func Newpasswordreset0040() *passwordreset0040 {
    return &passwordreset0040{}
}

func (e *passwordreset0040) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0040) Name() string { return "passwordreset0040" }
func (e *passwordreset0040) Timestamp() time.Time { return time.Now() }
