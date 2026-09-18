package grpc

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineServiceEnumerate(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.ServiceEnumerate()
	if err != nil {
		t.Errorf("ServiceEnumerate failed: %v", err)
	}
}

func TestEngineMethodDiscover(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.MethodDiscover()
	if err != nil {
		t.Errorf("MethodDiscover failed: %v", err)
	}
}

func TestEngineAuthBypass(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.AuthBypass()
	if err != nil {
		t.Errorf("AuthBypass failed: %v", err)
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
