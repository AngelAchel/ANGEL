package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type HMACSigner struct {
	key []byte
}

func NewHMACSigner(key []byte) *HMACSigner {
	return &HMACSigner{key: key}
}

func (s *HMACSigner) Sign(data string) string {
	mac := hmac.New(sha256.New, s.key)
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

func (s *HMACSigner) Verify(data, signature string) bool {
	expected := s.Sign(data)
	return hmac.Equal([]byte(expected), []byte(signature))
}

func (s *HMACSigner) Name() string         { return "HMACSigner" }
func (s *HMACSigner) Timestamp() time.Time { return time.Now() }
func (s *HMACSigner) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "hmac:running")
	return results, nil
}
