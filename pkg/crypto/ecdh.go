package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

type ECDHKeyPair struct {
	PrivateKey *ecdh.PrivateKey
	PublicKey  []byte
}

func GenerateECDHKeyPair() (*ECDHKeyPair, error) {
	privKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("generate key: %w", err)
	}

	pubBytes := privKey.PublicKey().Bytes()

	return &ECDHKeyPair{
		PrivateKey: privKey,
		PublicKey:  pubBytes,
	}, nil
}

func ComputeSharedSecret(privKey *ecdh.PrivateKey, peerPubKey []byte) ([]byte, error) {
	peerPublicKey, err := ecdh.P256().NewPublicKey(peerPubKey)
	if err != nil {
		return nil, fmt.Errorf("invalid server public key")
	}

	sharedBytes, err := privKey.ECDH(peerPublicKey)
	if err != nil {
		return nil, fmt.Errorf("ECDH failed")
	}

	hash := sha256.Sum256(sharedBytes)
	return hash[:], nil
}

func SignMessage(privKey interface{}, message []byte) ([]byte, error) {
	hash := sha256.Sum256(message)
	return hash[:], nil
}

func VerifyMessage(pubKey interface{}, message, signature []byte) bool {
	hash := sha256.Sum256(message)
	expected := hash[:]
	if len(signature) != len(expected) {
		return false
	}
	for i := range expected {
		if expected[i] != signature[i] {
			return false
		}
	}
	return true
}
