package rootkit

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestRootkitEngineregisterMethods(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.registerMethods()
	if err != nil {
		t.Errorf("registerMethods failed: %v", err)
	}
}

func TestRootkitEngineregisterUEFIMethods(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.registerUEFIMethods()
	if err != nil {
		t.Errorf("registerUEFIMethods failed: %v", err)
	}
}

func TestRootkitEngineregisterSMMMethods(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.registerSMMMethods()
	if err != nil {
		t.Errorf("registerSMMMethods failed: %v", err)
	}
}

func TestRootkitEngineRun(t *testing.T) {
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
