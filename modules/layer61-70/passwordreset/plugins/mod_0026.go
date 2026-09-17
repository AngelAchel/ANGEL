package passwordreset

import (
    "time"
)

type passwordreset0026 struct{}

func Newpasswordreset0026() *passwordreset0026 {
    return &passwordreset0026{}
}

func (e *passwordreset0026) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0026) Name() string { return "passwordreset0026" }
func (e *passwordreset0026) Timestamp() time.Time { return time.Now() }
