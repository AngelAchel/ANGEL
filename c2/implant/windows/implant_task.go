package implant

import (
	"time"
)

// Windows implant task execution
type WindowsTask struct{}

func NewWindowsTask() *WindowsTask {
	return &WindowsTask{}
}

func (t *WindowsTask) Execute() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "windows:task_executed")
	return results, nil
}

func (t *WindowsTask) Name() string { return "WindowsTask" }
func (t *WindowsTask) Timestamp() time.Time { return time.Now() }
