package crypto

import (
	"time"
)

type JavaObj struct{}

func NewJavaObj() *JavaObj {
	return &JavaObj{}
}

func (j *JavaObj) Process() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "java_obj:processed")
	return results, nil
}

func (j *JavaObj) Name() string         { return "JavaObj" }
func (j *JavaObj) Timestamp() time.Time { return time.Now() }
