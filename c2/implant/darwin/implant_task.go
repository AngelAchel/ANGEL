package implant

import (
	"time"
)

// Darwin implant task execution
type DarwinTask struct{}

func NewDarwinTask() *DarwinTask {
	return &DarwinTask{}
}

func (t *DarwinTask) Execute() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "darwin:task_executed")
	return results, nil
}

func (t *DarwinTask) Name() string         { return "DarwinTask" }
func (t *DarwinTask) Timestamp() time.Time { return time.Now() }
