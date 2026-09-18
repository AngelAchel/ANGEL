package arpdhcp

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineARPSpoof(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ARPSpoof()
	if err != nil {
		t.Errorf("ARPSpoof failed: %v", err)
	}
}

func TestEngineARPStorm(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ARPStorm()
	if err != nil {
		t.Errorf("ARPStorm failed: %v", err)
	}
}

func TestEngineDHCPRogue(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DHCPRogue()
	if err != nil {
		t.Errorf("DHCPRogue failed: %v", err)
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
