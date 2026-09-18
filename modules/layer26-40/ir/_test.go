package ir

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineBreachSimulate(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.BreachSimulate()
	if err != nil {
		t.Errorf("BreachSimulate failed: %v", err)
	}
}

func TestEngineContainment(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Containment()
	if err != nil {
		t.Errorf("Containment failed: %v", err)
	}
}

func TestEngineEradication(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Eradication()
	if err != nil {
		t.Errorf("Eradication failed: %v", err)
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
