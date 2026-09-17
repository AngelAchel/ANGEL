package sqli

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestSQLiEngineregisterDetectors(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.registerDetectors()
	if err != nil {
		t.Errorf("registerDetectors failed: %v", err)
	}
}

func TestSQLiEngineScan(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.Scan()
	if err != nil {
		t.Errorf("Scan failed: %v", err)
	}
}

func TestSQLiEngineDetectInjectionPoint(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.DetectInjectionPoint()
	if err != nil {
		t.Errorf("DetectInjectionPoint failed: %v", err)
	}
}

func TestSQLiEngineRun(t *testing.T) {
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
