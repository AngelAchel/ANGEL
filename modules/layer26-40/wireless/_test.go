package wireless

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineEvilTwin(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.EvilTwin()
	if err != nil {
		t.Errorf("EvilTwin failed: %v", err)
	}
}

func TestEngineDeauthAttack(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DeauthAttack()
	if err != nil {
		t.Errorf("DeauthAttack failed: %v", err)
	}
}

func TestEngineHandshakeCapture(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.HandshakeCapture()
	if err != nil {
		t.Errorf("HandshakeCapture failed: %v", err)
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
