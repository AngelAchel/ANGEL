package vlan

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineDoubleTagging(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DoubleTagging()
	if err != nil {
		t.Errorf("DoubleTagging failed: %v", err)
	}
}

func TestEngineSwitchSpoofing(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.SwitchSpoofing()
	if err != nil {
		t.Errorf("SwitchSpoofing failed: %v", err)
	}
}

func TestEngineVLANGrafting(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.VLANGrafting()
	if err != nil {
		t.Errorf("VLANGrafting failed: %v", err)
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
