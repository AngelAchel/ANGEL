package ipv6

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineRASpoof(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.RASpoof()
	if err != nil {
		t.Errorf("RASpoof failed: %v", err)
	}
}

func TestEnginebuildRAPacket(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.buildRAPacket()
	if err != nil {
		t.Errorf("buildRAPacket failed: %v", err)
	}
}

func TestEngineNSFlood(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.NSFlood()
	if err != nil {
		t.Errorf("NSFlood failed: %v", err)
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
