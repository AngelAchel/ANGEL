package physical

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineUSBDrop(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.USBDrop()
	if err != nil {
		t.Errorf("USBDrop failed: %v", err)
	}
}

func TestEngineBadgeClone(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.BadgeClone()
	if err != nil {
		t.Errorf("BadgeClone failed: %v", err)
	}
}

func TestEngineLockPicking(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.LockPicking()
	if err != nil {
		t.Errorf("LockPicking failed: %v", err)
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
