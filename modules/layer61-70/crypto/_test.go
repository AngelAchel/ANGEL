package crypto

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEnginePaddingOracle(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.PaddingOracle()
	if err != nil {
		t.Errorf("PaddingOracle failed: %v", err)
	}
}

func TestEngineWeakKeyDetect(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.WeakKeyDetect()
	if err != nil {
		t.Errorf("WeakKeyDetect failed: %v", err)
	}
}

func TestEngineHashLengthExt(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.HashLengthExt()
	if err != nil {
		t.Errorf("HashLengthExt failed: %v", err)
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
