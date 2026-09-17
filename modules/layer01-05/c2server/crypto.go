package c2server

import (
	"crypto/ecdh"
	"crypto/sha256"

	"github.com/angel-platform/angel/pkg/crypto"
)

type ServerCrypto struct {
	keyPair *crypto.ECDHKeyPair
	encKey  []byte
}

func NewServerCrypto() (*ServerCrypto, error) {
	keyPair, err := crypto.GenerateECDHKeyPair()
	if err != nil {
		return nil, err
	}
	encKey := crypto.GenerateAESKey()
	return &ServerCrypto{
		keyPair: keyPair,
		encKey:  encKey,
	}, nil
}

func NewServerCryptoFromKey(privKey *ecdh.PrivateKey) (*ServerCrypto, error) {
	pubBytes := privKey.PublicKey().Bytes()
	encKey := crypto.GenerateAESKey()
	return &ServerCrypto{
		keyPair: &crypto.ECDHKeyPair{
			PrivateKey: privKey,
			PublicKey:  pubBytes,
		},
		encKey: encKey,
	}, nil
}

func (s *ServerCrypto) PublicKey() *ecdh.PublicKey {
	if s.keyPair != nil && len(s.keyPair.PublicKey) > 0 {
		pubKey, err := ecdh.P256().NewPublicKey(s.keyPair.PublicKey)
		if err == nil {
			return pubKey
		}
	}
	return nil
}

func (s *ServerCrypto) EncryptPayload(data []byte) ([]byte, error) {
	return crypto.EncryptAESGCM(s.encKey, data)
}

func (s *ServerCrypto) DecryptPayload(data []byte) ([]byte, error) {
	return crypto.DecryptAESGCM(s.encKey, data)
}

func (s *ServerCrypto) IsEstablished() bool {
	return s.keyPair != nil && len(s.keyPair.PublicKey) > 0
}

func (s *ServerCrypto) SignData(data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	return hash[:], nil
}

func (s *ServerCrypto) VerifySignature(data []byte, sig []byte) bool {
	hash := sha256.Sum256(data)
	return string(hash[:]) == string(sig)
}
