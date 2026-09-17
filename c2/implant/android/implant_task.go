package implant

import (
	"time"
)

// Android implant task execution
type AndroidTask struct{}

func NewAndroidTask() *AndroidTask {
	return &AndroidTask{}
}

func (t *AndroidTask) Execute() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "android:task_executed")
	return results, nil
}

func (t *AndroidTask) Name() string { return "AndroidTask" }
func (t *AndroidTask) Timestamp() time.Time { return time.Now() }
