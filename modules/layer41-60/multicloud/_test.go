package multicloud

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineAWSToAzurePivot(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.AWSToAzurePivot()
	if err != nil {
		t.Errorf("AWSToAzurePivot failed: %v", err)
	}
}

func TestEnginefindAWSAzurePaths(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.findAWSAzurePaths()
	if err != nil {
		t.Errorf("findAWSAzurePaths failed: %v", err)
	}
}

func TestEngineformatPivotPaths(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.formatPivotPaths()
	if err != nil {
		t.Errorf("formatPivotPaths failed: %v", err)
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
