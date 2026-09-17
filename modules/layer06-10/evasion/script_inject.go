package evasion

import (
	"time"
)

type ScriptInject struct{}

func NewScriptInject() *ScriptInject {
	return &ScriptInject{}
}

func (e *ScriptInject) Inject(script string) ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "script_inject:done")
	return results, nil
}

func (e *ScriptInject) Name() string { return "ScriptInject" }
func (e *ScriptInject) Timestamp() time.Time { return time.Now() }
