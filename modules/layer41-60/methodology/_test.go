package methodology

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineReconPhase(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ReconPhase()
	if err != nil {
		t.Errorf("ReconPhase failed: %v", err)
	}
}

func TestEngineperformRecon(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.performRecon()
	if err != nil {
		t.Errorf("performRecon failed: %v", err)
	}
}

func TestEngineformatReconFindings(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.formatReconFindings()
	if err != nil {
		t.Errorf("formatReconFindings failed: %v", err)
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
