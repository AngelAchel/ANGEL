package mobile

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineKeychainDump(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.KeychainDump()
	if err != nil {
		t.Errorf("KeychainDump failed: %v", err)
	}
}

func TestEngineSSLPinningBypass(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.SSLPinningBypass()
	if err != nil {
		t.Errorf("SSLPinningBypass failed: %v", err)
	}
}

func TestEngineSharedPreferencesExtract(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.SharedPreferencesExtract()
	if err != nil {
		t.Errorf("SharedPreferencesExtract failed: %v", err)
	}
}

func TestEngineRun(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	result, err := gw.Run()
	if err != nil {
		t.Errorf("Run failed: %v", err)
	}
	if result == "" {
		t.Error("expected non-empty result")
	}
}
