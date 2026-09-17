package persistence

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestPersistenceEngineregisterPlatformMethods(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.registerPlatformMethods()
	if err != nil {
		t.Errorf("registerPlatformMethods failed: %v", err)
	}
}

func TestPersistenceEngineregisterWindowsMethods(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.registerWindowsMethods()
	if err != nil {
		t.Errorf("registerWindowsMethods failed: %v", err)
	}
}

func TestPersistenceEngineregisterLinuxMethods(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.registerLinuxMethods()
	if err != nil {
		t.Errorf("registerLinuxMethods failed: %v", err)
	}
}

func TestPersistenceEngineRun(t *testing.T) {
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
