package webmisc

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)

func TestEngineCachePoisoning(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.CachePoisoning()
	if err != nil {
		t.Errorf("CachePoisoning failed: %v", err)
	}
}

func TestEngineanalyzeCachePoisonVectors(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.analyzeCachePoisonVectors()
	if err != nil {
		t.Errorf("analyzeCachePoisonVectors failed: %v", err)
	}
}

func TestEngineformatPoisonMethods(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.formatPoisonMethods()
	if err != nil {
		t.Errorf("formatPoisonMethods failed: %v", err)
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
