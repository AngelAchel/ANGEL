package c2server

import (
	"time"
)

type Component struct{}

func NewComponent() *Component {
	return &Component{}
}

func (e *Component) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "component:done")
	return results, nil
}

func (e *Component) Name() string         { return "Component" }
func (e *Component) Timestamp() time.Time { return time.Now() }
