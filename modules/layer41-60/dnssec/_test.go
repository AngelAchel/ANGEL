package dnssec

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineNSECWalking(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.NSECWalking()
	if err != nil {
		t.Errorf("NSECWalking failed: %v", err)
	}
}

func TestEngineperformNSECWalk(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.performNSECWalk()
	if err != nil {
		t.Errorf("performNSECWalk failed: %v", err)
	}
}

func TestEnginesimulateNSECRecord(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.simulateNSECRecord()
	if err != nil {
		t.Errorf("simulateNSECRecord failed: %v", err)
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
