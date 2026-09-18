package credential

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestCredentialEngineregisterExtractors(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.registerExtractors()
	if err != nil {
		t.Errorf("registerExtractors failed: %v", err)
	}
}

func TestCredentialEngineHarvest(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Harvest()
	if err != nil {
		t.Errorf("Harvest failed: %v", err)
	}
}

func TestCredentialEngineharvestLSASS(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.harvestLSASS()
	if err != nil {
		t.Errorf("harvestLSASS failed: %v", err)
	}
}

func TestCredentialEngineRun(t *testing.T) {
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
