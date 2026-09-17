package evasion

import (
	"time"
)

type XPCmdshell struct{}

func NewXPCmdshell() *XPCmdshell {
	return &XPCmdshell{}
}

func (e *XPCmdshell) Execute(cmd string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "xp_cmdshell:done")
	return results, nil
}

func (e *XPCmdshell) Name() string         { return "XPCmdshell" }
func (e *XPCmdshell) Timestamp() time.Time { return time.Now() }
