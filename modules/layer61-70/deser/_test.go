package deser

import (
	"testing"

	"github.com/angel-platform/angel/gateway"
)


func TestEngineJavaDeserExploit(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.JavaDeserExploit()
	if err != nil {
		t.Errorf("JavaDeserExploit failed: %v", err)
	}
}

func TestEnginePythonPickle(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.PythonPickle()
	if err != nil {
		t.Errorf("PythonPickle failed: %v", err)
	}
}

func TestEnginePHPSerialize(t *testing.T) {
	gw := New(gateway.DefaultConfig())
	if gw == nil {
		t.Fatal("expected non-nil engine")
	}
	_, err := gw.PHPSerialize()
	if err != nil {
		t.Errorf("PHPSerialize failed: %v", err)
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
