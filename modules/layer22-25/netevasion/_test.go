package netevasion

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestNetEvasionEngineEvade(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Evade()
	if err != nil {
		t.Errorf("Evade failed: %v", err)
	}
}

func TestNetEvasionEngineFullEvasion(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.FullEvasion()
	if err != nil {
		t.Errorf("FullEvasion failed: %v", err)
	}
}

func TestNetEvasionEngineevadeIPRotation(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.evadeIPRotation()
	if err != nil {
		t.Errorf("evadeIPRotation failed: %v", err)
	}
}

func TestNetEvasionEngineRun(t *testing.T) {
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
