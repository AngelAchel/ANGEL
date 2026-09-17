package evasion

import (
	"time"
)

type Commands struct{}

func NewCommands() *Commands {
	return &Commands{}
}

func (e *Commands) Execute(cmd string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "commands:executed")
	return results, nil
}

func (e *Commands) Name() string { return "Commands" }
func (e *Commands) Category() EvasionCategory { return CategorySyscall }
func (e *Commands) Timestamp() time.Time { return time.Now() }
