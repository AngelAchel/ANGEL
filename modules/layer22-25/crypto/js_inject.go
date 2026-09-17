package crypto

import (
	"time"
)

type JSInject struct{}

func NewJSInject() *JSInject {
	return &JSInject{}
}

func (j *JSInject) Inject() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "js_inject:done")
	return results, nil
}

func (j *JSInject) Name() string         { return "JSInject" }
func (j *JSInject) Timestamp() time.Time { return time.Now() }
