package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
)

type ServerCrypto struct {
	privateKey *ecdh.PrivateKey
	publicKey  *ecdh.PublicKey
	aesKey     []byte
}

func NewServerCrypto() (*ServerCrypto, error) {
	privateKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	pubKey := privateKey.PublicKey()
	return &ServerCrypto{
		privateKey: privateKey,
		publicKey:  pubKey,
		aesKey:     generateAESKey(),
	}, nil
}

func generateAESKey() []byte {
	key := make([]byte, 32)
	rand.Read(key)
	return key
}
