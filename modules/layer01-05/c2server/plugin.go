package c2server

import (
	"time"
)

type Plugin struct{}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (e *Plugin) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "plugin:done")
	return results, nil
}

func (e *Plugin) Name() string         { return "Plugin" }
func (e *Plugin) Timestamp() time.Time { return time.Now() }
