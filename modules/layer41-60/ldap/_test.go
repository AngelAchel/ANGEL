package ldap

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineFilterInjection(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.FilterInjection()
	if err != nil {
		t.Errorf("FilterInjection failed: %v", err)
	}
}

func TestEnginebuildInjectionFilter(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.buildInjectionFilter()
	if err != nil {
		t.Errorf("buildInjectionFilter failed: %v", err)
	}
}

func TestEngineNullBind(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.NullBind()
	if err != nil {
		t.Errorf("NullBind failed: %v", err)
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
