package implant

import (
	"time"
)

type ImplantTask struct{}

func NewImplantTask() *ImplantTask {
	return &ImplantTask{}
}

func (i *ImplantTask) Task() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "implant_task:done")
	return results, nil
}

func (i *ImplantTask) Name() string         { return "ImplantTask" }
func (i *ImplantTask) Timestamp() time.Time { return time.Now() }
