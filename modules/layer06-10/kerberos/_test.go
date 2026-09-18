package kerberos

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestKerberosEngineGoldenTicket(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.GoldenTicket()
	if err != nil {
		t.Errorf("GoldenTicket failed: %v", err)
	}
}

func TestKerberosEngineSilverTicket(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.SilverTicket()
	if err != nil {
		t.Errorf("SilverTicket failed: %v", err)
	}
}

func TestKerberosEngineDiamondTicket(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DiamondTicket()
	if err != nil {
		t.Errorf("DiamondTicket failed: %v", err)
	}
}

func TestKerberosEngineRun(t *testing.T) {
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
