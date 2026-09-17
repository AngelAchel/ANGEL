package implant

import (
	"time"
)

type ImplantCrypto struct{}

func NewImplantCrypto() *ImplantCrypto {
	return &ImplantCrypto{}
}

func (i *ImplantCrypto) Encrypt() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "implant_crypto:done")
	return results, nil
}

func (i *ImplantCrypto) Name() string { return "ImplantCrypto" }
func (i *ImplantCrypto) Timestamp() time.Time { return time.Now() }
