package implant

import (
	"time"
)

type ImplantResult struct{}

func NewImplantResult() *ImplantResult {
	return &ImplantResult{}
}

func (i *ImplantResult) Result() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "implant_result:done")
	return results, nil
}

func (i *ImplantResult) Name() string         { return "ImplantResult" }
func (i *ImplantResult) Timestamp() time.Time { return time.Now() }
