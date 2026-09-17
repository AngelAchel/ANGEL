package implant

import (
	"time"
)

// Darwin implant registration
type DarwinRegister struct{}

func NewDarwinRegister() *DarwinRegister {
	return &DarwinRegister{}
}

func (r *DarwinRegister) Register() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "darwin:registered")
	return results, nil
}

func (r *DarwinRegister) Name() string { return "DarwinRegister" }
func (r *DarwinRegister) Timestamp() time.Time { return time.Now() }
