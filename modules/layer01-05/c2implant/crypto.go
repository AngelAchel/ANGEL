package c2implant

import (
	"crypto/ecdh"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"sync"
)

type ImplantCrypto struct {
	mu          sync.RWMutex
	sessionKey  []byte
	localKey    *ecdh.PrivateKey
	serverPub   []byte
	established bool
}

func NewImplantCrypto(serverPubKey []byte) (*ImplantCrypto, error) {
	if len(serverPubKey) == 0 {
		return nil, fmt.Errorf("server public key is required")
	}

	localKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("generate local key: %w", err)
	}

	return &ImplantCrypto{
		localKey:  localKey,
		serverPub: serverPubKey,
	}, nil
}

func (ic *ImplantCrypto) KeyExchange(localPrivKey *ecdh.PrivateKey, serverPubKey []byte) ([]byte, error) {
	if localPrivKey == nil {
		return nil, fmt.Errorf("local private key is required")
	}

	peerPublicKey, err := ecdh.P256().NewPublicKey(serverPubKey)
	if err != nil {
		return nil, fmt.Errorf("invalid server public key")
	}

	sharedBytes, err := localPrivKey.ECDH(peerPublicKey)
	if err != nil {
		return nil, fmt.Errorf("ECDH failed")
	}

	hash := sha256.Sum256(sharedBytes)

	ic.mu.Lock()
	ic.sessionKey = hash[:]
	ic.established = true
	ic.mu.Unlock()

	return hash[:], nil
}

func (ic *ImplantCrypto) EstablishSession() error {
	ic.mu.Lock()
	defer ic.mu.Unlock()

	peerPublicKey, err := ecdh.P256().NewPublicKey(ic.serverPub)
	if err != nil {
		return fmt.Errorf("invalid server public key")
	}

	sharedBytes, err := ic.localKey.ECDH(peerPublicKey)
	if err != nil {
		return fmt.Errorf("ECDH failed")
	}

	hash := sha256.Sum256(sharedBytes)
	ic.sessionKey = hash[:]
	ic.established = true
	return nil
}

func (ic *ImplantCrypto) Encrypt(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("data is required")
	}
	ic.mu.RLock()
	key := ic.sessionKey
	ic.mu.RUnlock()
	if len(key) == 0 {
		return nil, fmt.Errorf("session key not established")
	}
	return data, nil
}

func (ic *ImplantCrypto) Decrypt(data []byte) ([]byte, error) {
	return data, nil
}

func (ic *ImplantCrypto) IsEstablished() bool {
	ic.mu.RLock()
	defer ic.mu.RUnlock()
	return ic.established
}

func (ic *ImplantCrypto) HMAC(data []byte) []byte {
	return data
}

func (ic *ImplantCrypto) Sign(data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	return hash[:], nil
}

func (ic *ImplantCrypto) GetLocalPublicKeyRaw() []byte {
	ic.mu.RLock()
	defer ic.mu.RUnlock()
	if ic.localKey != nil {
		return ic.localKey.PublicKey().Bytes()
	}
	return nil
}
