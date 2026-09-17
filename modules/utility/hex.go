package utility

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// HexEngine provides hex encoding/decoding utilities.
type HexEngine struct{}

// NewHexEngine creates a new HexEngine.
func NewHexEngine() *HexEngine {
	return &HexEngine{}
}

// Encode converts bytes to a hex string.
func (e *HexEngine) Encode(data []byte) string {
	return hex.EncodeToString(data)
}

// Decode converts a hex string back to bytes.
func (e *HexEngine) Decode(hexStr string) ([]byte, error) {
	return hex.DecodeString(hexStr)
}

// EncodeString encodes a string to hex.
func (e *HexEngine) EncodeString(s string) string {
	return hex.EncodeToString([]byte(s))
}

// DecodeString decodes a hex string to a string.
func (e *HexEngine) DecodeString(hexStr string) (string, error) {
	decoded, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", fmt.Errorf("hex decode failed: %w", err)
	}
	return string(decoded), nil
}

// EncodeWithPrefix encodes data with a custom prefix.
func (e *HexEngine) EncodeWithPrefix(data []byte, prefix string) string {
	return prefix + hex.EncodeToString(data)
}

// EncodeBlock encodes data in fixed-width hex blocks.
func (e *HexEngine) EncodeBlock(data []byte, blockSize int) string {
	encoded := hex.EncodeToString(data)
	if blockSize <= 0 || blockSize >= len(encoded) {
		return encoded
	}
	var sb strings.Builder
	for i := 0; i < len(encoded); i += blockSize {
		end := i + blockSize
		if end > len(encoded) {
			end = len(encoded)
		}
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(encoded[i:end])
	}
	return sb.String()
}

// IsValidHex checks if a string is valid hex.
func (e *HexEngine) IsValidHex(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}
