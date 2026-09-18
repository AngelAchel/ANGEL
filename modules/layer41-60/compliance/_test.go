package compliance

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestDataDeletionRecordPCICompliance(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.PCICompliance()
	if err != nil {
		t.Errorf("PCICompliance failed: %v", err)
	}
}

func TestDataDeletionRecordcheckPCIControls(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.checkPCIControls()
	if err != nil {
		t.Errorf("checkPCIControls failed: %v", err)
	}
}

func TestDataDeletionRecordGDPRCheck(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.GDPRCheck()
	if err != nil {
		t.Errorf("GDPRCheck failed: %v", err)
	}
}

func TestDataDeletionRecordRun(t *testing.T) {
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
