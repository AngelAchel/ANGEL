package implant

import (
	"time"
)

// Android implant main entry point
type AndroidMain struct{}

func NewAndroidMain() *AndroidMain {
	return &AndroidMain{}
}

func (m *AndroidMain) Run() error {
	return nil
}

func (m *AndroidMain) Name() string { return "AndroidMain" }
func (m *AndroidMain) Timestamp() time.Time { return time.Now() }
