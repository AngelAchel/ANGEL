package passwordreset

import (
    "time"
)

type passwordreset0028 struct{}

func Newpasswordreset0028() *passwordreset0028 {
    return &passwordreset0028{}
}

func (e *passwordreset0028) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0028) Name() string { return "passwordreset0028" }
func (e *passwordreset0028) Timestamp() time.Time { return time.Now() }
