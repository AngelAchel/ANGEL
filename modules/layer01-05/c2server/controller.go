package c2server

import (
	"time"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (e *Controller) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "controller:done")
	return results, nil
}

func (e *Controller) Name() string         { return "Controller" }
func (e *Controller) Timestamp() time.Time { return time.Now() }
