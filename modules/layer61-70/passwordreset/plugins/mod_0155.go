package passwordreset

import (
    "time"
)

type passwordreset0155 struct{}

func Newpasswordreset0155() *passwordreset0155 {
    return &passwordreset0155{}
}

func (e *passwordreset0155) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0155) Name() string { return "passwordreset0155" }
func (e *passwordreset0155) Timestamp() time.Time { return time.Now() }
