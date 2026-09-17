package passwordreset

import (
    "time"
)

type passwordreset0182 struct{}

func Newpasswordreset0182() *passwordreset0182 {
    return &passwordreset0182{}
}

func (e *passwordreset0182) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0182) Name() string { return "passwordreset0182" }
func (e *passwordreset0182) Timestamp() time.Time { return time.Now() }
