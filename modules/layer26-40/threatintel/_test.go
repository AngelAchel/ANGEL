package threatintel

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineIOCGeneration(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.IOCGeneration()
	if err != nil {
		t.Errorf("IOCGeneration failed: %v", err)
	}
}

func TestEngineMITREATTACKMapping(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.MITREATTACKMapping()
	if err != nil {
		t.Errorf("MITREATTACKMapping failed: %v", err)
	}
}

func TestEngineThreatFeed(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ThreatFeed()
	if err != nil {
		t.Errorf("ThreatFeed failed: %v", err)
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
