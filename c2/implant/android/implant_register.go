package implant

import (
	"time"
)

// Android implant registration
type AndroidRegister struct{}

func NewAndroidRegister() *AndroidRegister {
	return &AndroidRegister{}
}

func (r *AndroidRegister) Register() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "android:registered")
	return results, nil
}

func (r *AndroidRegister) Name() string         { return "AndroidRegister" }
func (r *AndroidRegister) Timestamp() time.Time { return time.Now() }
