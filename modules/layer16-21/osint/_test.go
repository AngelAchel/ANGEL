package osint

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestOSINTEngineRecon(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Recon()
	if err != nil {
		t.Errorf("Recon failed: %v", err)
	}
}

func TestOSINTEngineFullRecon(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.FullRecon()
	if err != nil {
		t.Errorf("FullRecon failed: %v", err)
	}
}

func TestOSINTEngineGetStats(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.GetStats()
	if err != nil {
		t.Errorf("GetStats failed: %v", err)
	}
}

func TestOSINTEngineRun(t *testing.T) {
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
