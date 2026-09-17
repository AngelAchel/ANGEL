package memcorrupt

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineBufferOverflowDetect(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.BufferOverflowDetect()
	if err != nil {
		t.Errorf("BufferOverflowDetect failed: %v", err)
	}
}

func TestEngineUseAfterFree(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.UseAfterFree()
	if err != nil {
		t.Errorf("UseAfterFree failed: %v", err)
	}
}

func TestEngineHeapSpray(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.HeapSpray()
	if err != nil {
		t.Errorf("HeapSpray failed: %v", err)
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
