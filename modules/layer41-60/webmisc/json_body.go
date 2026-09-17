package webmisc

import (
	"time"
)

type JSONBody struct{}

func NewJSONBody() *JSONBody {
	return &JSONBody{}
}

func (j *JSONBody) Inject() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "json_body:done")
	return results, nil
}

func (j *JSONBody) Name() string { return "JSONBody" }
func (j *JSONBody) Timestamp() time.Time { return time.Now() }
