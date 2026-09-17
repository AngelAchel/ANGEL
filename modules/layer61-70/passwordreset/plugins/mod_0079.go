package passwordreset

import (
    "time"
)

type passwordreset0079 struct{}

func Newpasswordreset0079() *passwordreset0079 {
    return &passwordreset0079{}
}

func (e *passwordreset0079) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0079) Name() string { return "passwordreset0079" }
func (e *passwordreset0079) Timestamp() time.Time { return time.Now() }
