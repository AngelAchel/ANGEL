package passwordreset

import (
    "time"
)

type passwordreset0129 struct{}

func Newpasswordreset0129() *passwordreset0129 {
    return &passwordreset0129{}
}

func (e *passwordreset0129) Run() ([]string, error) {
    results := make([]string, 0, 1)
    results = append(results, "passwordreset:done")
    return results, nil
}

func (e *passwordreset0129) Name() string { return "passwordreset0129" }
func (e *passwordreset0129) Timestamp() time.Time { return time.Now() }
